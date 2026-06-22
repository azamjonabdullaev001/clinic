# Полное руководство по развёртыванию на новом сервере

Этот документ описывает установку проекта «с нуля» на новый VPS/сервер, включая
настройку домена, Nginx и HTTPS (Let's Encrypt). Весь проект работает в Docker —
ничего, кроме Docker и (внешнего) Nginx, на сервере ставить не нужно.

---

## 1. Архитектура и порты

Проект состоит из трёх контейнеров (см. `docker-compose.yml`):

| Сервис     | Технология          | Порт на хосте | Назначение                          |
|------------|---------------------|---------------|-------------------------------------|
| `db`       | PostgreSQL 16       | `5435 → 5432` | База данных                         |
| `backend`  | Go (Gin) + GORM     | `8082 → 8080` | REST API + WebSocket (`/api/...`)   |
| `frontend` | Vue 3 + Nginx       | `3002 → 80`   | Сайт + проксирование `/api` на backend |

Внутренний Nginx фронтенда (`frontend/nginx.conf`) уже проксирует `/api/` и
`/uploads/` на контейнер `backend`, поэтому **снаружи достаточно открыть только
порт `3002`**. Внешний (системный) Nginx на сервере проксирует домен → `127.0.0.1:3002`
и завершает HTTPS.

```
Интернет → :443 (системный Nginx, SSL) → 127.0.0.1:3002 (контейнер frontend)
                                              ├── / (статика Vue)
                                              └── /api/, /uploads/ → backend:8080
```

---

## 2. Требования к серверу

- Ubuntu 22.04 / 24.04 (или Debian 12) с root-доступом
- Минимум 1 vCPU / 2 GB RAM (рекомендуется 2 vCPU / 4 GB)
- Открытые порты `80` и `443`
- Доменное имя, A-запись которого указывает на IP сервера

---

## 3. Установка Docker

```bash
# обновляем систему
apt update && apt upgrade -y

# ставим Docker + Docker Compose plugin официальным скриптом
curl -fsSL https://get.docker.com | sh

# проверяем
docker --version
docker compose version
```

---

## 4. Получение кода

```bash
# каталог по умолчанию, который ожидает deploy.sh
cd /root
git clone <URL_ВАШЕГО_РЕПОЗИТОРИЯ> clinic
cd clinic
```

> Если переносите проект архивом — просто распакуйте его в `/root/clinic`.

---

## 5. Файл `.env` (обязательно)

В корне проекта создайте `.env` на основе `.env.example`:

```bash
cp .env.example .env
nano .env
```

Заполните **все** значения собственными (не оставляйте дефолтные пароли!):

```env
POSTGRES_DB=clinic
POSTGRES_USER=clinic
POSTGRES_PASSWORD=<надёжный_пароль_БД>
JWT_SECRET=<длинная_случайная_строка_минимум_32_символа>
ADMIN_PHONE=998901475130
ADMIN_PASSWORD=<пароль_админа>
TELEGRAM_BOT_TOKEN=<токен_бота_или_оставьте_пустым>
TELEGRAM_CHAT_ID=<chat_id_или_оставьте_пустым>
```

Подсказки:
- Сгенерировать `JWT_SECRET`: `openssl rand -hex 32`
- `ADMIN_PHONE` / `ADMIN_PASSWORD` — данные для входа в админ-панель. Админ
  создаётся автоматически при первом запуске, если в БД ещё нет ни одного админа.
- Telegram-переменные нужны только для уведомлений; их можно оставить пустыми.

---

## 6. Запуск проекта

```bash
cd /root/clinic
docker compose up -d --build
```

Проверка состояния:

```bash
docker compose ps          # все три сервиса должны быть Up (db — healthy)
docker compose logs -f backend   # логи backend (Ctrl+C для выхода)
```

Быстрая проверка локально на сервере:

```bash
curl -I http://127.0.0.1:3002          # должен вернуть 200 / отдать index.html
curl http://127.0.0.1:3002/api/products # JSON-список товаров (онлайн-доступные)
```

База данных хранится в Docker-томе `pgdata`, загруженные изображения — в томе
`uploads`. Они переживают перезапуск и пересборку контейнеров.

---

## 7. Домен (DNS)

В панели вашего регистратора/DNS создайте записи:

| Тип | Имя              | Значение            |
|-----|------------------|---------------------|
| A   | `@` (домен)      | IP вашего сервера   |
| A   | `www`            | IP вашего сервера   |

Дождитесь распространения DNS (`dig +short ваш-домен` должен вернуть IP сервера).

---

## 8. Системный Nginx + HTTPS

### 8.1. Установка

```bash
apt install -y nginx certbot python3-certbot-nginx
```

### 8.2. Стартовый конфиг (только HTTP, для выпуска сертификата)

Создайте `/etc/nginx/sites-available/clinic.conf`, заменив `ВАШ_ДОМЕН`:

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name ВАШ_ДОМЕН www.ВАШ_ДОМЕН;

    location / {
        proxy_pass http://127.0.0.1:3002;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # WebSocket для живого обновления склада/заказов (/api/ws/stock)
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 3600s;

        client_max_body_size 10M;   # загрузка чеков/фото препаратов
    }
}
```

Активируйте и проверьте:

```bash
ln -s /etc/nginx/sites-available/clinic.conf /etc/nginx/sites-enabled/
rm -f /etc/nginx/sites-enabled/default
nginx -t && systemctl reload nginx
```

### 8.3. Выпуск сертификата Let's Encrypt

```bash
certbot --nginx -d ВАШ_ДОМЕН -d www.ВАШ_ДОМЕН
```

Certbot сам перепишет конфиг: добавит блок `listen 443 ssl`, пути к сертификатам
и редирект с HTTP на HTTPS (как в готовом примере `nginx/doctor-jalilov.uz.conf`).
Сертификаты лягут в `/etc/letsencrypt/live/ВАШ_ДОМЕН/`.

> **Важно:** убедитесь, что в итоговом `server`-блоке для `location /` остались
> заголовки `Upgrade`/`Connection "upgrade"` — без них перестанет работать
> WebSocket живого обновления. Если certbot их убрал, верните их вручную и
> выполните `nginx -t && systemctl reload nginx`.

### 8.4. Автопродление сертификата

Certbot ставит systemd-таймер автоматически. Проверьте:

```bash
systemctl status certbot.timer
certbot renew --dry-run
```

После этого сайт доступен по `https://ВАШ_ДОМЕН`.

---

## 9. Демоданные (необязательно)

Чтобы наполнить базу тестовыми заказами/остатками, есть идемпотентный сид
(подробности — в `scripts/README.md`):

```bash
bash scripts/run_seed_full.sh
```

Скрипт сначала делает бэкап текущей БД в `clinic_backup_<дата>.sql`, затем
перезаписывает данные. **На боевых данных не запускайте.**

---

## 10. Доступ в панели

- **Сайт:** `https://ВАШ_ДОМЕН`
- **Админ-панель:** `https://ВАШ_ДОМЕН/admin/login`
  - Телефон/пароль — из `.env` (`ADMIN_PHONE` / `ADMIN_PASSWORD`)
- Работники (пункт выдачи / медсестра / менеджер / врач) создаются в админ-панели.

Не забудьте положить фото врача: `frontend/public/images/doctor.jpg`
(пересобрать фронтенд после замены: `docker compose up -d --build frontend`).

---

## 11. Обновление кода

```bash
cd /root/clinic
git pull origin main
docker compose up -d --build
```

В репозитории есть готовый `deploy.sh`, делающий то же самое:

```bash
bash deploy.sh
```

---

## 12. Резервное копирование и восстановление

**Бэкап базы:**

```bash
docker exec clinic-db-1 pg_dump -U clinic -d clinic > backup_$(date +%F).sql
```

**Бэкап загруженных файлов (том `uploads`):**

```bash
docker run --rm -v clinic_uploads:/data -v $(pwd):/backup alpine \
  tar czf /backup/uploads_$(date +%F).tar.gz -C /data .
```

**Восстановление базы:**

```bash
cat backup.sql | docker exec -i clinic-db-1 psql -U clinic -d clinic
```

> Имя контейнера БД (`clinic-db-1`) зависит от имени каталога проекта. Проверьте
> через `docker compose ps`.

---

## 13. Частые проблемы

| Симптом | Причина / решение |
|---------|-------------------|
| `502 Bad Gateway` | Контейнер `frontend` не запущен или порт `3002` занят. `docker compose ps`, `docker compose logs frontend`. |
| API возвращает CORS-ошибку | Запросы должны идти через тот же домен (`/api/...`), а не напрямую на `:8082`. Системный Nginx уже это обеспечивает. |
| Не работает живое обновление склада | В конфиге Nginx для `location /` нет заголовков `Upgrade`/`Connection`. См. п. 8.2. |
| `JWT_SECRET ... must be set` в логах backend | Не заполнен `.env` или контейнер запущен без него. Пересоберите: `docker compose up -d --build`. |
| Не загружаются фото (413) | `client_max_body_size 10M;` должен быть в обоих Nginx (внешний — п. 8.2; внутренний уже настроен). |
| Сертификат не выпускается | Проверьте, что DNS указывает на сервер и порт 80 открыт в фаерволе. |

---

## 14. Фаервол (рекомендуется)

```bash
ufw allow OpenSSH
ufw allow 'Nginx Full'   # 80 + 443
ufw enable
```

Порты `3002`, `8082`, `5435` наружу открывать **не нужно** — они используются
только локально между Nginx и контейнерами. При желании можно ограничить их в
`docker-compose.yml`, привязав к `127.0.0.1` (например `"127.0.0.1:3002:80"`).

---

## 15. Краткая шпаргалка (всё по шагам)

```bash
# 1. Docker
curl -fsSL https://get.docker.com | sh

# 2. Код
cd /root && git clone <REPO> clinic && cd clinic

# 3. Настройки
cp .env.example .env && nano .env

# 4. Запуск
docker compose up -d --build

# 5. Nginx + SSL
apt install -y nginx certbot python3-certbot-nginx
# создать /etc/nginx/sites-available/clinic.conf (см. п. 8.2), активировать
certbot --nginx -d ВАШ_ДОМЕН -d www.ВАШ_ДОМЕН

# Готово: https://ВАШ_ДОМЕН
```
