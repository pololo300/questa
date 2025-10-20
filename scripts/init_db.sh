#!/bin/bash
set -e

# === conf ===
DB_NAME="questa"
DB_USER="questauser"
DB_PASS="questapass"
DB_DIR="/srv/questa/data"
PG_VERSION="15"

# === mkdir ===
mkdir -p "$DB_DIR"
chown -R postgres:postgres "$DB_DIR"

echo "Initializing dabase PostgreSQL in $DB_DIR..."

# === launch cluster ===
sudo -u postgres /usr/lib/postgresql/${PG_VERSION}/bin/initdb -D "$DB_DIR"

# === launch postgress on background ===
sudo -u postgres /usr/lib/postgresql/${PG_VERSION}/bin/pg_ctl -D "$DB_DIR" -l "$DB_DIR/logfile" start
sleep 3

# === setup user and database ===
sudo -u postgres psql <<EOF
DO
\$do\$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles WHERE rolname = '${DB_USER}'
   ) THEN
      CREATE ROLE ${DB_USER} LOGIN PASSWORD '${DB_PASS}';
   END IF;
END
\$do\$;

CREATE DATABASE ${DB_NAME} OWNER ${DB_USER};
GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} TO ${DB_USER};
EOF

# === show info ===
echo "✅ database created correctly:"
echo "   - Host: 192.169.1.32"
echo "   - Purt: 5432"
echo "   - Name: ${DB_NAME}"
echo "   - User: ${DB_USER}"
echo "   - Password: ${DB_PASS}"
echo "   - Data: ${DB_DIR}"
