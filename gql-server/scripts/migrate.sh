MIGRATIONS_DIR="./database/migrations"

cd ${MIGRATIONS_DIR}

goose sqlite3 ../todos.db up