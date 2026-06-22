#!/bin/bash
# =============================================================================
#  backup_old_server.sh  —  Eski VPI serverdan to'liq backup olish
# =============================================================================
#  Ishlatish (ESK serverda root yoki sudo bilan):
#
#    chmod +x backup_old_server.sh
#    ./backup_old_server.sh
#
#  Natija:  /root/clinic_backup_YYYYMMDD_HHMMSS.tar.gz
#           (ichida: clinic_db.sql  +  uploads/ papkasi)
#
#  Keyin bu faylni yangi serverga ko'chiring:
#    scp /root/clinic_backup_*.tar.gz root@YANGI_SERVER_IP:/root/
# =============================================================================
set -euo pipefail

# ── Ranglar ──────────────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()    { echo -e "${CYAN}[INFO]${NC}  $*"; }
success() { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC}  $*"; }
error()   { echo -e "${RED}[ERR]${NC}   $*" >&2; exit 1; }

# ── Sozlamalar ────────────────────────────────────────────────────────────────
CLINIC_DIR="${CLINIC_DIR:-/root/clinic}"
BACKUP_DIR="${BACKUP_DIR:-/root}"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_NAME="clinic_backup_${TIMESTAMP}"
WORK_DIR="/tmp/${BACKUP_NAME}"

# ── Boshlash ──────────────────────────────────────────────────────────────────
echo ""
echo "======================================================"
echo "   CLINIC BACKUP — eski VPI server"
echo "   $(date '+%Y-%m-%d %H:%M:%S')"
echo "======================================================"
echo ""

# ── 1. Docker va konteyner tekshirish ─────────────────────────────────────────
info "Docker konteynerlari tekshirilmoqda..."
if ! command -v docker &>/dev/null; then
    error "Docker topilmadi. Docker o'rnatilganligini tekshiring."
fi

DB_CONTAINER=$(docker ps --format '{{.Names}}' | grep -E 'db|postgres' | head -1)
if [ -z "$DB_CONTAINER" ]; then
    # Try finding by image
    DB_CONTAINER=$(docker ps --format '{{.Names}}\t{{.Image}}' | grep postgres | awk '{print $1}' | head -1)
fi
if [ -z "$DB_CONTAINER" ]; then
    warn "Ishlab turgan DB konteyner topilmadi. docker ps chiqishi:"
    docker ps --format 'table {{.Names}}\t{{.Image}}\t{{.Status}}'
    error "DB konteynerni qo'lda ko'rsating: DB_CONTAINER=clinic-db-1 ./backup_old_server.sh"
fi
success "DB konteyner: ${DB_CONTAINER}"

# ── 2. .env dan sozlamalar o'qish ─────────────────────────────────────────────
info ".env fayl o'qilmoqda..."
ENV_FILE="${CLINIC_DIR}/.env"
if [ ! -f "$ENV_FILE" ]; then
    # Try one level up
    ENV_FILE="$(dirname "$CLINIC_DIR")/.env"
fi
if [ ! -f "$ENV_FILE" ]; then
    warn ".env fayl topilmadi, default qiymatlar ishlatiladi."
    POSTGRES_USER="clinic"
    POSTGRES_DB="clinic"
else
    # shellcheck disable=SC1090
    source "$ENV_FILE"
    POSTGRES_USER="${POSTGRES_USER:-clinic}"
    POSTGRES_DB="${POSTGRES_DB:-clinic}"
fi
success "DB: ${POSTGRES_DB}  User: ${POSTGRES_USER}"

# ── 3. Ishchi papka tayyorlash ────────────────────────────────────────────────
info "Ishchi papka tayyorlanmoqda: ${WORK_DIR}"
mkdir -p "${WORK_DIR}"

# ── 4. PostgreSQL dump ────────────────────────────────────────────────────────
info "PostgreSQL dump olinmoqda..."
SQL_FILE="${WORK_DIR}/clinic_db.sql"

docker exec "$DB_CONTAINER" pg_dump \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB" \
    --no-password \
    --verbose \
    --format=plain \
    --encoding=UTF8 \
    > "$SQL_FILE" 2>/tmp/pgdump_log.txt

if [ $? -ne 0 ]; then
    warn "pg_dump logi:"
    cat /tmp/pgdump_log.txt
    error "pg_dump muvaffaqiyatsiz tugadi."
fi

SQL_SIZE=$(du -sh "$SQL_FILE" | cut -f1)
success "SQL dump tayyor: ${SQL_FILE} (${SQL_SIZE})"

# Jadvallar sonini tekshirish
TABLE_COUNT=$(grep -c "^CREATE TABLE" "$SQL_FILE" 2>/dev/null || echo "0")
info "Dump ichida jadvalar soni: ${TABLE_COUNT}"

# ── 5. Uploads volume backup ──────────────────────────────────────────────────
info "Uploads fayllari backup qilinmoqda..."
UPLOADS_DIR="${WORK_DIR}/uploads"
mkdir -p "$UPLOADS_DIR"

# Docker volume orqali ko'chirish
BACKEND_CONTAINER=$(docker ps --format '{{.Names}}' | grep -E 'backend|app' | head -1)

if [ -n "$BACKEND_CONTAINER" ]; then
    info "Backend konteyneridan ko'chirilmoqda: ${BACKEND_CONTAINER}"
    docker cp "${BACKEND_CONTAINER}:/app/uploads/." "$UPLOADS_DIR/" 2>/dev/null || true
    UPLOAD_COUNT=$(find "$UPLOADS_DIR" -type f | wc -l)
    success "Uploads ko'chirildi: ${UPLOAD_COUNT} ta fayl"
else
    warn "Backend konteyner topilmadi, volume orqali urinilmoqda..."
    # Find the volume mount point
    VOLUME_PATH=$(docker volume inspect clinic_uploads 2>/dev/null | \
        python3 -c "import sys,json; d=json.load(sys.stdin); print(d[0]['Mountpoint'])" 2>/dev/null || \
        docker volume inspect "$(basename "$CLINIC_DIR")_uploads" 2>/dev/null | \
        python3 -c "import sys,json; d=json.load(sys.stdin); print(d[0]['Mountpoint'])" 2>/dev/null || \
        echo "")
    if [ -n "$VOLUME_PATH" ] && [ -d "$VOLUME_PATH" ]; then
        cp -r "${VOLUME_PATH}/." "$UPLOADS_DIR/" 2>/dev/null || true
        UPLOAD_COUNT=$(find "$UPLOADS_DIR" -type f | wc -l)
        success "Uploads ko'chirildi (volume orqali): ${UPLOAD_COUNT} ta fayl"
    else
        warn "Uploads volume topilmadi, bo'sh papka saqlanadi."
    fi
fi

# ── 6. Meta-info saqlash ──────────────────────────────────────────────────────
info "Meta-info yozilmoqda..."
cat > "${WORK_DIR}/backup_info.txt" << METAEOF
BACKUP_DATE=$(date '+%Y-%m-%d %H:%M:%S')
BACKUP_TIMESTAMP=${TIMESTAMP}
HOSTNAME=$(hostname)
POSTGRES_USER=${POSTGRES_USER}
POSTGRES_DB=${POSTGRES_DB}
DB_CONTAINER=${DB_CONTAINER}
SQL_SIZE=${SQL_SIZE}
UPLOAD_FILES=${UPLOAD_COUNT:-0}
TABLE_COUNT=${TABLE_COUNT}
METAEOF
success "backup_info.txt yozildi"

# ── 7. Siqish ─────────────────────────────────────────────────────────────────
FINAL_ARCHIVE="${BACKUP_DIR}/${BACKUP_NAME}.tar.gz"
info "Arxiv yaratilmoqda: ${FINAL_ARCHIVE}"
tar -czf "$FINAL_ARCHIVE" -C /tmp "$BACKUP_NAME"
ARCHIVE_SIZE=$(du -sh "$FINAL_ARCHIVE" | cut -f1)
success "Arxiv tayyor: ${FINAL_ARCHIVE} (${ARCHIVE_SIZE})"

# ── 8. Tozalash ───────────────────────────────────────────────────────────────
rm -rf "$WORK_DIR"
info "Vaqtinchalik fayllar o'chirildi."

# ── Yakuniy xabar ─────────────────────────────────────────────────────────────
echo ""
echo "======================================================"
echo -e "  ${GREEN}BACKUP MUVAFFAQIYATLI YAKUNLANDI!${NC}"
echo "======================================================"
echo ""
echo "  Arxiv:   ${FINAL_ARCHIVE}"
echo "  O'lcham: ${ARCHIVE_SIZE}"
echo ""
echo "  Yangi serverga ko'chirish uchun:"
echo ""
echo "  scp ${FINAL_ARCHIVE} root@YANGI_SERVER_IP:/root/"
echo ""
echo "  Keyin yangi serverda:"
echo "  ./restore_new_server.sh /root/${BACKUP_NAME}.tar.gz"
echo ""
