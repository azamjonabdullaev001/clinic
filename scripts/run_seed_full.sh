#!/usr/bin/env bash
# Rebuild all demo data on the server. Run from the repo root ON THE SERVER.
# Usage:  bash scripts/run_seed_full.sh
set -euo pipefail

DB_CONTAINER="${DB_CONTAINER:-clinic-db-1}"
DB_USER="${DB_USER:-clinic}"
DB_NAME="${DB_NAME:-clinic}"
SQL_FILE="$(dirname "$0")/seed_full.sql"

echo ">> Backing up current DB ..."
docker exec "$DB_CONTAINER" pg_dump -U "$DB_USER" -d "$DB_NAME" \
  > "clinic_backup_$(date +%Y%m%d_%H%M%S).sql"

echo ">> Copying seed into container ..."
docker cp "$SQL_FILE" "$DB_CONTAINER":/seed_full.sql

echo ">> Running seed (this takes ~30-60s) ..."
docker exec -i "$DB_CONTAINER" psql -U "$DB_USER" -d "$DB_NAME" -v ON_ERROR_STOP=1 -f /seed_full.sql

echo ">> Done."
