#!/bin/bash
export PGPASSWORD=$(cat /run/secrets/postgres_password)
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'