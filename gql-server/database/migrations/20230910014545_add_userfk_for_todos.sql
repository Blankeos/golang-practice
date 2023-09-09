-- +goose Up
-- +goose StatementBegin
PRAGMA foreign_keys=off;

-- 1. Create Default User
INSERT INTO users (id, name) VALUES (80, 'Default User');

-- 2. Rename the original table to old
ALTER TABLE todos RENAME TO _todos_old;

-- 3. Create the new todos table with a foreign key for users
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT,
    done BOOLEAN DEFAULT 0,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 4. Paste back old todos into the new todos + the default foreign key.
INSERT INTO todos (id, text, done, user_id)
SELECT id, text, done, 80 FROM _todos_old;

-- 5. Drop the old todos table
DROP TABLE _todos_old;

PRAGMA foreign_keys=on;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
PRAGMA foreign_keys=off;

-- 1. Remove the user_id column from the todos table
CREATE TABLE todos_backup (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    text TEXT,
    done BOOLEAN DEFAULT 0
);

INSERT INTO todos_backup (id, text, done)
SELECT id, text, done FROM todos;

DROP TABLE todos;

ALTER TABLE todos_backup RENAME TO todos;

-- 2. Delete the Default User
DELETE FROM users WHERE id = 80;
-- +goose StatementEnd
