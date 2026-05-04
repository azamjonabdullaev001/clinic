# Hair Clinic — Веб-приложение

## Запуск

```bash
docker-compose up --build
```

- **Сайт**: http://localhost:3002
- **API**: http://localhost:8082

## Администрирование домена

Для сервера VPS нужно настроить внешнюю Nginx-конфигурацию для домена `doctor-jalilov.uz`.

- HTTP должен перенаправляться на HTTPS
- HTTPS должен проксировать трафик на локальный контейнер фронтенда: `http://127.0.0.1:3002`
- Сертификаты Let's Encrypt ожидаются в `/etc/letsencrypt/live/doctor-jalilov.uz/`

Файл примера конфигурации находится в `nginx/doctor-jalilov.uz.conf`.

## Админ-панель

- **URL**: http://localhost:3000/admin/login
- **Телефон**: +998 90 147 51 30
- **Пароль**: 12345678

## Важно

Поместите фото доктора в файл:
```
frontend/public/images/doctor.jpg
```

## Стек

- **Frontend**: Vue 3, Vite, Tailwind CSS, Pinia, Vue Router
- **Backend**: Go (Gin), GORM, JWT
- **Database**: PostgreSQL 16
- **Infra**: Docker, Docker Compose
