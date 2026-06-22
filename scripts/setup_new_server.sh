#!/bin/bash
# =============================================================================
#  setup_new_server.sh  —  Yangi Ubuntu/Debian serverda muhit tayyorlash
# =============================================================================
#  Restore qilishdan OLDIN yangi serverda bir marta ishlatiladi.
#  Bu script: Docker, Git, Nginx, Certbot o'rnatadi va loyihani klonlaydi.
#
#  Ishlatish (yangi serverda root sifatida):
#    curl -fsSL https://raw.githubusercontent.com/azamjonabdullaev001/clinic/main/scripts/setup_new_server.sh | bash
#  Yoki:
#    chmod +x setup_new_server.sh && ./setup_new_server.sh
# =============================================================================
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()    { echo -e "${CYAN}[INFO]${NC}  $*"; }
success() { echo -e "${GREEN}[OK]${NC}    $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC}  $*"; }
error()   { echo -e "${RED}[ERR]${NC}   $*" >&2; exit 1; }

CLINIC_DIR="${CLINIC_DIR:-/root/clinic}"
DOMAIN="${DOMAIN:-doctor-jalilov.uz}"

echo ""
echo "======================================================"
echo "   YANGI SERVER SETUP"
echo "   $(date '+%Y-%m-%d %H:%M:%S')"
echo "======================================================"
echo ""

# ── Root tekshirish ───────────────────────────────────────────────────────────
if [ "$EUID" -ne 0 ]; then
    error "Bu script root yoki sudo bilan ishlatilishi kerak!"
fi

# ── OS tekshirish ─────────────────────────────────────────────────────────────
if [ -f /etc/os-release ]; then
    # shellcheck disable=SC1091
    source /etc/os-release
    info "OS: ${PRETTY_NAME}"
fi

# ── 1. Tizim yangilash ────────────────────────────────────────────────────────
info "Tizim paketlari yangilanmoqda..."
apt-get update -qq
apt-get upgrade -y -qq
apt-get install -y -qq \
    ca-certificates curl gnupg lsb-release \
    git nginx certbot python3-certbot-nginx \
    ufw htop net-tools
success "Asosiy paketlar o'rnatildi."

# ── 2. Docker o'rnatish ───────────────────────────────────────────────────────
if command -v docker &>/dev/null; then
    success "Docker allaqachon o'rnatilgan: $(docker --version)"
else
    info "Docker o'rnatilmoqda..."
    install -m 0755 -d /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg \
        | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    chmod a+r /etc/apt/keyrings/docker.gpg

    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
        | tee /etc/apt/sources.list.d/docker.list > /dev/null

    apt-get update -qq
    apt-get install -y -qq docker-ce docker-ce-cli containerd.io docker-compose-plugin
    systemctl enable docker
    systemctl start docker
    success "Docker o'rnatildi: $(docker --version)"
fi

# ── 3. Git loyiha klonlash ────────────────────────────────────────────────────
if [ -d "$CLINIC_DIR/.git" ]; then
    info "Loyiha allaqachon klonlangan. git pull qilinmoqda..."
    git -C "$CLINIC_DIR" pull origin main || warn "git pull muvaffaqiyatsiz (lokal o'zgarishlar bo'lishi mumkin)"
else
    info "Loyiha klonlanmoqda → ${CLINIC_DIR}"
    git clone https://github.com/azamjonabdullaev001/clinic.git "$CLINIC_DIR"
    success "Loyiha klonlandi."
fi

# ── 4. .env fayl yaratish ─────────────────────────────────────────────────────
ENV_FILE="${CLINIC_DIR}/.env"
if [ ! -f "$ENV_FILE" ]; then
    info ".env fayl yaratilmoqda..."

    # Tasodifiy kalit generatsiya
    JWT_SECRET=$(openssl rand -hex 32)
    DB_PASS=$(openssl rand -base64 20 | tr -dc 'A-Za-z0-9' | head -c 20)

    cat > "$ENV_FILE" << ENVEOF
# ── Database ──────────────────────────────────────────
POSTGRES_DB=clinic
POSTGRES_USER=clinic
POSTGRES_PASSWORD=${DB_PASS}

# ── Backend ───────────────────────────────────────────
JWT_SECRET=${JWT_SECRET}

# ── Admin akkaunt (birinchi login) ────────────────────
ADMIN_PHONE=+998901234567
ADMIN_PASSWORD=admin123

# ── Telegram bildirishnomalar (ixtiyoriy) ─────────────
TELEGRAM_BOT_TOKEN=
TELEGRAM_CHAT_ID=
ENVEOF

    chmod 600 "$ENV_FILE"
    warn ".env fayl yaratildi: ${ENV_FILE}"
    warn "Admin phone va parolni HOZIROQ o'zgartiring!"
    cat "$ENV_FILE"
else
    success ".env fayl allaqachon mavjud."
fi

# ── 5. Nginx konfiguratsiya ───────────────────────────────────────────────────
NGINX_CONF="/etc/nginx/sites-available/${DOMAIN}"
if [ ! -f "$NGINX_CONF" ]; then
    info "Nginx konfiguratsiya yaratilmoqda..."
    cat > "$NGINX_CONF" << NGINXEOF
server {
    listen 80;
    listen [::]:80;
    server_name ${DOMAIN} www.${DOMAIN};

    # Frontend
    location / {
        proxy_pass http://127.0.0.1:3002;
        proxy_http_version 1.1;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
        proxy_set_header Connection "";
        proxy_read_timeout 60s;
        client_max_body_size 50M;
    }

    # Backend API
    location /api/ {
        proxy_pass http://127.0.0.1:8082;
        proxy_http_version 1.1;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 120s;
        client_max_body_size 50M;
    }

    # WebSocket
    location /ws {
        proxy_pass http://127.0.0.1:8082;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host \$host;
        proxy_read_timeout 3600s;
    }
}
NGINXEOF

    ln -sf "$NGINX_CONF" /etc/nginx/sites-enabled/
    rm -f /etc/nginx/sites-enabled/default
    nginx -t && systemctl reload nginx
    success "Nginx sozlandi."
else
    success "Nginx konfiguratsiya allaqachon mavjud."
fi

# ── 6. Firewall sozlash ───────────────────────────────────────────────────────
info "Firewall sozlanmoqda..."
ufw allow OpenSSH
ufw allow 'Nginx Full'
ufw --force enable
success "UFW faol: SSH + HTTP/HTTPS ruxsat berildi."

# ── Yakuniy xabar ─────────────────────────────────────────────────────────────
SERVER_IP=$(curl -s ifconfig.me 2>/dev/null || hostname -I | awk '{print $1}')

echo ""
echo "======================================================"
echo -e "  ${GREEN}SETUP YAKUNLANDI!${NC}"
echo "======================================================"
echo ""
echo "  Server IP: ${SERVER_IP}"
echo "  Loyiha:    ${CLINIC_DIR}"
echo "  Domain:    ${DOMAIN}"
echo ""
echo "  Keyingi qadamlar:"
echo ""
echo "  1. .env ni tahrirlang (admin parol va boshqa):"
echo "     nano ${CLINIC_DIR}/.env"
echo ""
echo "  2. Backup faylni ko'chiring:"
echo "     scp /path/to/clinic_backup_*.tar.gz root@${SERVER_IP}:/root/"
echo ""
echo "  3. Restore qiling:"
echo "     chmod +x ${CLINIC_DIR}/scripts/restore_new_server.sh"
echo "     ${CLINIC_DIR}/scripts/restore_new_server.sh /root/clinic_backup_*.tar.gz"
echo ""
echo "  4. DNS ni yangi serverga yo'naltiring:"
echo "     ${DOMAIN}  A  ${SERVER_IP}"
echo ""
echo "  5. SSL sertifikat (DNS o'zgarishidan keyin):"
echo "     certbot --nginx -d ${DOMAIN} -d www.${DOMAIN}"
echo ""
