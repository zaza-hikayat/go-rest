#!/bin/sh

. ./.env

echo "development:
  dialect: postgres
  database: ${DB_NAME}
  user: ${DB_USER}
  password: ${DB_PASSWORD}
  host: ${DB_HOST}
  pool: 5" > database.yml