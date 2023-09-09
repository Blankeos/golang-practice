-- name: GetTodo :one
SELECT * from todos
WHERE id = ? LIMIT 1;

-- name: GetTodos :many
SELECT * FROM todos
ORDER BY text;

-- name: CreateTodo :one
INSERT INTO todos (
    text
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;