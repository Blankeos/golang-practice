-- name: GetTodo :one
SELECT * from todos
WHERE id = ? LIMIT 1;

-- name: GetTodosWithUser :many
SELECT todos.*, users.*
FROM todos
LEFT JOIN users ON todos.user_id = users.id
ORDER BY text;

-- name: GetTodos :many
SELECT * FROM todos
ORDER BY text;

-- name: CreateTodo :one
INSERT INTO todos (
    text,
    user_id
) VALUES (
    ?1,
    ?2
)
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;