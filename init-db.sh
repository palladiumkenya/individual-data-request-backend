#!/bin/bash
export PGPASSWORD=${POSTGRES_PASSWORD}
psql -U idr -d idr -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'
