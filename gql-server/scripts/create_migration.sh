MIGRATIONS_DIR="./database/migrations"
MIGRATION_NAME=$1

cd ${MIGRATIONS_DIR}

goose sqlite3 ../todos.db create ${MIGRATION_NAME} sql