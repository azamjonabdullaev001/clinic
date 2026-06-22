#!/bin/bash
# =============================================================================
#  restore_new_server.sh  —  Yangi VPI serverga clinic loyihasini ko'chirish
# =============================================================================
#  Ishlatish (YANGI serverda root yoki sudo bilan):
#
#    chmod +x restore_new_server.sh
#    ./restore_new_server.sh /root/clinic_backup_YYYYMMDD_HHMMSS.tar.gz
#
#  Oldindan talab qilinadigan narsalar:
#    - Docker + Docker Compose o'rnatilgan bo'lsin
#    - /root/clinic papkasida loyiha klonlangan bo'lsin (git clone)
#    - /root/clinic/.env fayl to'ldirilgan bo'lsin (quyida namuna bor)
#
#  .env namunasi:
#    POSTGRES_DB=clinic
#    POSTGRES_USER=clinic
#    POSTGRES_PASSWORD=KUCHLI_PAROL
#    JWT_SECRET=UZUN_TASODIFIY_KALIT
#    ADMIN_PHONE=+998901234567
#    ADMIN_PASSWORD=ADMIN_PAROLI
#    TELEGRAM_BOT_TOKEN=
#    TELEGRAM_CHAT_ID=
# =============================================================================
set -euo pipefail

# ── Ranglar ──────────────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()    { echo -e "${CYAN}[INFO]${NC}  $*"; }
success() { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC}  $*"; }
error()   { echo -e "${RED}[ERR]${NC}   $*" >&2; exit 1; }
prompt()  { echo -e "${YELLOW}[???]${NC}  $*"; }

# ── Argument tekshirish ───────────────────────────────────────────────────────
if [ $# -lt 1 ]; then
    error "Backup fayl yo'li ko'rsatilmagan!\n\n  Ishlatish: $0 /root/clinic_backup_YYYYMMDD_HHMMSS.tar.gz"
fi

BACKUP_ARCHIVE="$1"
if [ ! -f "$BACKUP_ARCHIVE" ]; then
    error "Fayl topilmadi: ${BACKUP_ARCHIVE}"
fi

CLINIC_DIR="${CLINIC_DIR:-/root/clinic}"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
WORK_DIR="/tmp/clinic_restore_${TIMESTAMP}"

# ── Boshlash ──────────────────────────────────────────────────────────────────
echo ""
echo "======================================================"
echo "   CLINIC RESTORE — yangi VPI server"
echo "   $(date '+%Y-%m-%d %H:%M:%S')"
echo "======================================================"
echo ""

# ── 0. Talablar tekshirish ────────────────────────────────────────────────────
info "Tizim talablari tekshirilmoqda..."

command -v docker   &>/dev/null || error "Docker o'rnatilmagan!"
command -v git      &>/dev/null || error "Git o'rnatilmagan!"

COMPOSE_CMD=""
if docker compose version &>/dev/null 2>&1; then
    COMPOSE_CMD="docker compose"
elif command -v docker-compose &>/dev/null; then
    COMPOSE_CMD="docker-compose"
else
    error "Docker Compose topilmadi! Avval o'rnating: apt install docker-compose-plugin"
fi
success "Docker Compose: ${COMPOSE_CMD}"

# ── 1. Loyiha papkasi tekshirish ──────────────────────────────────────────────
info "Loyiha papkasi tekshirilmoqda: ${CLINIC_DIR}"
if [ ! -d "$CLINIC_DIR" ]; then
    warn "Loyiha papkasi topilmadi. Git clone qilinmoqda..."
    git clone https://github.com/azamjonabdullaev001/clinic.git "$CLINIC_DIR" \
        || error "Git clone muvaffaqiyatsiz. Qo'lda klonlang va qayta ishga tushiring."
fi
success "Loyiha papkasi: ${CLINIC_DIR}"

# ── 2. .env tekshirish ────────────────────────────────────────────────────────
ENV_FILE="${CLINIC_DIR}/.env"
if [ ! -f "$ENV_FILE" ]; then
    warn ".env fayl topilmadi!"
    echo ""
    echo "  Quyidagi faylni yarating: ${ENV_FILE}"
    echo ""
    cat << 'ENVEXAMPLE'
# ── .env fayl namunasi ──────────────────────────────
POSTGRES_DB=clinic
POSTGRES_USER=clinic
POSTGRES_PASSWORD=KUCHLI_PAROL_BU_YERGA

JWT_SECRET=KAMIDA_32_BELGILI_TASODIFIY_KALIT
ADMIN_PHONE=+998901234567
ADMIN_PASSWORD=ADMIN_PAROLI

TELEGRAM_BOT_TOKEN=
TELEGRAM_CHAT_ID=
# ────────────────────────────────────────────────────
ENVEXAMPLE
    echo ""
    error ".env fayl yaratib, qayta ishga tushiring."
fi

# .env o'qish
# shellcheck disable=SC1090
source "$ENV_FILE"
POSTGRES_USER="${POSTGRES_USER:-clinic}"
POSTGRES_DB="${POSTGRES_DB:-clinic}"
POSTGRES_PASSWORD="${POSTGRES_PASSWORD:-}"

if [ -z "$POSTGRES_PASSWORD" ]; then
    error ".env da POSTGRES_PASSWORD bo'sh! To'ldiring."
fi
success ".env o'qildi. DB: ${POSTGRES_DB}, User: ${POSTGRES_USER}"

# ── 3. Backup arxivini ochish ─────────────────────────────────────────────────
info "Backup arxivi ochilmoqda: ${BACKUP_ARCHIVE}"
mkdir -p "$WORK_DIR"
tar -xzf "$BACKUP_ARCHIVE" -C "$WORK_DIR"

# Ichki papkani topish
INNER_DIR=$(find "$WORK_DIR" -maxdepth 2 -name "clinic_db.sql" -exec dirname {} \; | head -1)
if [ -z "$INNER_DIR" ]; then
    error "Backup arxivida clinic_db.sql topilmadi. Arxiv yaroqsiz."
fi
success "Backup ochildi: ${INNER_DIR}"

SQL_FILE="${INNER_DIR}/clinic_db.sql"
UPLOADS_BACKUP="${INNER_DIR}/uploads"

# Backup ma'lumotlari
if [ -f "${INNER_DIR}/backup_info.txt" ]; then
    info "Backup ma'lumotlari:"
    cat "${INNER_DIR}/backup_info.txt" | sed 's/^/    /'
fi

# ── 4. Docker Compose ishga tushirish ─────────────────────────────────────────
info "Docker konteynerlari ishga tushirilmoqda..."
cd "$CLINIC_DIR"

# Faqat DB ni avval tushiramiz (migration uchun)
$COMPOSE_CMD up -d db
info "DB tayyor bo'lishi kutilmoqda (max 60 soniya)..."

READY=0
for i in $(seq 1 30); do
    if $COMPOSE_CMD exec -T db pg_isready -U "$POSTGRES_USER" &>/dev/null 2>&1; then
        READY=1
        break
    fi
    sleep 2
done

if [ $READY -eq 0 ]; then
    error "DB 60 soniyada tayyor bo'lmadi. Loglarni tekshiring: $COMPOSE_CMD logs db"
fi
success "PostgreSQL tayyor."

# ── 5. DB mavjudligini tekshirish ─────────────────────────────────────────────
DB_CONTAINER=$($COMPOSE_CMD ps -q db)
DB_CONTAINER_NAME=$($COMPOSE_CMD ps db --format '{{.Name}}' 2>/dev/null || docker ps --filter "id=$DB_CONTAINER" --format '{{.Names}}')

info "DB konteyner: ${DB_CONTAINER_NAME}"

# Database mavjud bo'lsa, tozalash tavsiya etiladi
EXISTS=$($COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -lqt 2>/dev/null | grep -c "$POSTGRES_DB" || echo "0")
if [ "$EXISTS" -gt 0 ]; then
    warn "DB '${POSTGRES_DB}' allaqachon mavjud."
    prompt "Mavjud DB ni tozalab restore qilamizmi? (yes/no)"
    read -r CONFIRM
    if [ "$CONFIRM" != "yes" ]; then
        error "Bekor qilindi. Qo'lda tozalang yoki yangi DB nomi bilan qayta urinib ko'ring."
    fi
    info "Mavjud DB o'chirilmoqda va qayta yaratilmoqda..."
    $COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -c "DROP DATABASE IF EXISTS ${POSTGRES_DB};"
    $COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -c "CREATE DATABASE ${POSTGRES_DB} ENCODING 'UTF8' LC_COLLATE='en_US.utf8' LC_CTYPE='en_US.utf8' TEMPLATE template0;"
else
    info "DB '${POSTGRES_DB}' mavjud emas, avtomatik yaratiladi..."
    $COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -c \
        "SELECT 1 FROM pg_database WHERE datname='${POSTGRES_DB}';" 2>/dev/null | grep -q 1 || \
        $COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -c \
        "CREATE DATABASE ${POSTGRES_DB} ENCODING 'UTF8' TEMPLATE template0;" 2>/dev/null || true
fi
success "DB tayyor."

# ── 6. SQL dump restore ───────────────────────────────────────────────────────
info "SQL dump restore qilinmoqda (bu bir necha daqiqa olishi mumkin)..."
SQL_SIZE=$(du -sh "$SQL_FILE" | cut -f1)
info "Dump o'lchami: ${SQL_SIZE}"

$COMPOSE_CMD exec -T db psql \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB" \
    < "$SQL_FILE" \
    2>/tmp/restore_log.txt

if [ $? -ne 0 ]; then
    warn "Restore logi (xatolar):"
    grep -i "error\|fatal" /tmp/restore_log.txt | head -20 || true
    # Ba'zi warning-lar normal — jadvallar soni tekshiramiz
fi

# Jadvallar soni tekshirish
TABLE_COUNT=$($COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -t -c \
    "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema='public';" 2>/dev/null | tr -d ' ')
success "Restore yakunlandi. Jadvalar soni: ${TABLE_COUNT}"

# Asosiy jadvallar tekshirish
for TABLE in users products orders workers doctors settings; do
    COUNT=$($COMPOSE_CMD exec -T db psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -t -c \
        "SELECT COUNT(*) FROM ${TABLE};" 2>/dev/null | tr -d ' ' || echo "?")
    info "  ${TABLE}: ${COUNT} ta yozuv"
done

# ── 7. Uploads restore ────────────────────────────────────────────────────────
info "Upload fayllar restore qilinmoqda..."

# Backend konteynerni ishga tushirish (uploads volume kerak)
$COMPOSE_CMD up -d backend
sleep 3

BACKEND_CONTAINER=$($COMPOSE_CMD ps backend --format '{{.Name}}' 2>/dev/null || \
    docker ps --filter "name=backend" --format '{{.Names}}' | head -1)

if [ -n "$BACKEND_CONTAINER" ] && [ -d "$UPLOADS_BACKUP" ]; then
    UPLOAD_COUNT=$(find "$UPLOADS_BACKUP" -type f | wc -l)
    if [ "$UPLOAD_COUNT" -gt 0 ]; then
        docker cp "${UPLOADS_BACKUP}/." "${BACKEND_CONTAINER}:/app/uploads/"
        success "Uploads ko'chirildi: ${UPLOAD_COUNT} ta fayl → /app/uploads/"
    else
        warn "Backup ichida upload fayllar topilmadi (bo'sh papka)."
    fi
else
    warn "Backend konteyner yoki uploads backup topilmadi, o'tkazib yuborildi."
fi

# ── 8. Barcha konteynerlarni ishga tushirish ──────────────────────────────────
info "Barcha konteynerlar ishga tushirilmoqda (build bilan)..."
$COMPOSE_CMD up -d --build

info "Konteynerlar tayyor bo'lishi kutilmoqda..."
sleep 10

# ── 9. Status tekshirish ──────────────────────────────────────────────────────
echo ""
info "Konteynerlar holati:"
$COMPOSE_CMD ps

# ── 10. Tozalash ──────────────────────────────────────────────────────────────
rm -rf "$WORK_DIR"
info "Vaqtinchalik fayllar o'chirildi."

# ── Yakuniy xabar ─────────────────────────────────────────────────────────────
echo ""
echo "======================================================"
echo -e "  ${GREEN}RESTORE MUVAFFAQIYATLI YAKUNLANDI!${NC}"
echo "======================================================"
echo ""
echo "  Tekshirish uchun:"
echo ""
echo "  Backend:   curl http://localhost:8082/api/health"
echo "  Frontend:  curl http://localhost:3002"
echo "  DB loglar: $COMPOSE_CMD logs db"
echo ""
echo "  SSL sertifikat yangilash (Certbot):"
echo "  certbot --nginx -d doctor-jalilov.uz -d www.doctor-jalilov.uz"
echo ""
