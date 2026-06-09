#!/bin/bash
# Yangi kod olish va saytni yangilash
set -e
cd /root/clinic
echo '=== git pull ==='
git pull origin main
echo '=== docker rebuild ==='
docker compose up -d --build
echo '=== status ==='
docker compose ps
echo '=== DONE ==='
