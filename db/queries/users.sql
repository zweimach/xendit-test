-- name: CreateUser :one
INSERT INTO users (
    login,
    name,
    avatar_url
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserByLogin :one
SELECT * FROM users
WHERE login = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListUsersByOrganizationLogin :many
SELECT u.* FROM users AS u
INNER JOIN memberships AS m
ON m.user_id = u.id
INNER JOIN organizations AS o
ON m.organization_id = o.id
WHERE o.login = $1
    AND o.deleted_at IS NULL
    AND u.deleted_at IS NULL
    AND m.deleted_at IS NULL
ORDER BY u.followers DESC
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET name = $2, avatar_url = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = now()
WHERE id = $1;
