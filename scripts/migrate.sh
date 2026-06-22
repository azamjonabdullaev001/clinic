#!/bin/bash
# =============================================================================
#  migrate.sh  —  Eski VPI serverdan yangi serverga TO'LIQ ko'chirish
# =============================================================================
#  Bu script UCHINCHI MASHINA yoki LOCAL kompyuterdan ishlatiladi
#  (ikkala serverga SSH kirish imkoni bo'lgan joydan).
#
#  Ishlatish:
#    chmod +x migrate.sh
#    ./migrate.sh
#
#  Yoki argumentlar bilan:
#    OLD_HOST=old.server.ip NEW_HOST=new.server.ip \
#    OLD_USER=root NEW_USER=root \
#    ./migrate.sh
#
#  Agar siz ESK serverning o'zida ishlasangiz:
#    SKIP_SCP=true NEW_HOST=new.server.ip ./migrate.sh
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()    { echo -e "${CYAN}[INFO]${NC}  $*"; }
success() { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC}  $*"; }
error()   { echo -e "${RED}[ERR]${NC}   $*" >&2; exit 1; }

# ── Server manzillari ─────────────────────────────────────────────────────────
OLD_HOST="${OLD_HOST:-}"
NEW_HOST="${NEW_HOST:-}"
OLD_USER="${OLD_USER:-root}"
NEW_USER="${NEW_USER:-root}"
OLD_PORT="${OLD_PORT:-22}"
NEW_PORT="${NEW_PORT:-22}"
SKIP_SCP="${SKIP_SCP:-false}"

echo ""
echo "======================================================"
echo "   CLINIC MIGRATION — eski → yangi server"
echo "   $(date '+%Y-%m-%d %H:%M:%S')"
echo "======================================================"
echo ""

# ── Serverlar so'rash ─────────────────────────────────────────────────────────
if [ -z "$OLD_HOST" ]; then
    read -rp "  Eski server IP/hostname: " OLD_HOST
fi
if [ -z "$NEW_HOST" ]; then
    read -rp "  Yangi server IP/hostname: " NEW_HOST
fi

SSH_OLD="${OLD_USER}@${OLD_HOST}"
SSH_NEW="${NEW_USER}@${NEW_HOST}"

info "Eski server:  ${SSH_OLD}:${OLD_PORT}"
info "Yangi server: ${SSH_NEW}:${NEW_PORT}"
echo ""

# ── SSH ulanish tekshirish ────────────────────────────────────────────────────
info "SSH ulanishlar tekshirilmoqda..."

if ! ssh -p "$OLD_PORT" -o ConnectTimeout=10 -o BatchMode=yes "$SSH_OLD" "echo ok" &>/dev/null; then
    error "Eski serverga SSH ulanib bo'lmadi: ${SSH_OLD}\nSSH kalitni tekshiring: ssh-copy-id -p ${OLD_PORT} ${SSH_OLD}"
fi
success "Eski server SSH: OK"

if ! ssh -p "$NEW_PORT" -o ConnectTimeout=10 -o BatchMode=yes "$SSH_NEW" "echo ok" &>/dev/null; then
    error "Yangi serverga SSH ulanib bo'lmadi: ${SSH_NEW}\nSSH kalitni tekshiring: ssh-copy-id -p ${NEW_PORT} ${SSH_NEW}"
fi
success "Yangi server SSH: OK"

# ── 1. Eski serverda backup olish ─────────────────────────────────────────────
info "Eski serverga backup scripti ko'chirilmoqda..."
scp -P "$OLD_PORT" "$(dirname "$0")/backup_old_server.sh" "${SSH_OLD}:/tmp/backup_old_server.sh"

info "Eski serverda backup ishga tushirilmoqda..."
BACKUP_FILE=$(ssh -p "$OLD_PORT" "$SSH_OLD" \
    "chmod +x /tmp/backup_old_server.sh && /tmp/backup_old_server.sh 2>&1 | tee /tmp/backup.log; \
     grep 'Arxiv:' /tmp/backup.log | awk '{print \$NF}'" | tail -1)

if [ -z "$BACKUP_FILE" ]; then
    # Try to find the latest backup manually
    BACKUP_FILE=$(ssh -p "$OLD_PORT" "$SSH_OLD" \
        "ls -t /root/clinic_backup_*.tar.gz 2>/dev/null | head -1" || echo "")
fi

if [ -z "$BACKUP_FILE" ]; then
    error "Backup fayli topilmadi! Eski serverda qo'lda tekshiring:\n  ssh ${SSH_OLD}\n  ls /root/clinic_backup_*.tar.gz"
fi
success "Backup tayyor: ${BACKUP_FILE}"

# ── 2. Backup faylni yangi serverga ko'chirish ────────────────────────────────
if [ "$SKIP_SCP" != "true" ]; then
    BACKUP_FILENAME=$(basename "$BACKUP_FILE")
    info "Backup yangi serverga ko'chirilmoqda (bu uzoq vaqt olishi mumkin)..."
    info "Manba:  ${SSH_OLD}:${BACKUP_FILE}"
    info "Maqsad: ${SSH_NEW}:/root/${BACKUP_FILENAME}"

    # rsync ishlatamiz (uzilsa davom etadi)
    if command -v rsync &>/dev/null; then
        rsync -avz --progress \
            -e "ssh -p ${OLD_PORT}" \
            "${SSH_OLD}:${BACKUP_FILE}" \
            "/tmp/${BACKUP_FILENAME}"
        rsync -avz --progress \
            -e "ssh -p ${NEW_PORT}" \
            "/tmp/${BACKUP_FILENAME}" \
            "${SSH_NEW}:/root/${BACKUP_FILENAME}"
        rm -f "/tmp/${BACKUP_FILENAME}"
    else
        # Direkt server-to-server ko'chirish (eski serverda SCP bilan)
        ssh -p "$OLD_PORT" "$SSH_OLD" \
            "scp -P ${NEW_PORT} ${BACKUP_FILE} ${SSH_NEW}:/root/${BACKUP_FILENAME}"
    fi
    success "Backup yangi serverga ko'chirildi: /root/${BACKUP_FILENAME}"
else
    BACKUP_FILENAME=$(basename "$BACKUP_FILE")
    info "SKIP_SCP=true, fayl allaqachon yangi serverda deb qabul qilinmoqda."
fi

# ── 3. Restore scriptini yangi serverga ko'chirish ────────────────────────────
info "Restore scripti yangi serverga ko'chirilmoqda..."
scp -P "$NEW_PORT" "$(dirname "$0")/restore_new_server.sh" "${SSH_NEW}:/tmp/restore_new_server.sh"

# ── 4. Yangi serverda restore ─────────────────────────────────────────────────
info "Yangi serverda restore ishga tushirilmoqda..."
ssh -p "$NEW_PORT" "$SSH_NEW" \
    "chmod +x /tmp/restore_new_server.sh && \
     CLINIC_DIR=/root/clinic /tmp/restore_new_server.sh /root/${BACKUP_FILENAME}"

success "Restore yakunlandi!"

# ── 5. Tekshirish ─────────────────────────────────────────────────────────────
info "Yangi server holatini tekshirish..."
ssh -p "$NEW_PORT" "$SSH_NEW" \
    "cd /root/clinic && docker compose ps 2>/dev/null || docker-compose ps"

echo ""
echo "======================================================"
echo -e "  ${GREEN}MIGRATSIYA MUVAFFAQIYATLI YAKUNLANDI!${NC}"
echo "======================================================"
echo ""
echo "  Yangi server tekshirish:"
echo "  ssh ${SSH_NEW}"
echo "  curl http://localhost:8082/api/health"
echo ""
echo "  DNS ni yangi serverga yo'naltiring:"
echo "  doctor-jalilov.uz  →  ${NEW_HOST}"
echo ""
echo "  SSL sertifikat:"
echo "  certbot --nginx -d doctor-jalilov.uz -d www.doctor-jalilov.uz"
echo ""
