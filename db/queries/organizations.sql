-- name: CreateOrganization :one
INSERT INTO organizations (
    login,
    name
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetOrganizationByID :one
SELECT * FROM organizations
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetOrganizationByLogin :one
SELECT * FROM organizations
WHERE login = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: ListOrganizations :many
SELECT * FROM organizations
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOrganization :one
UPDATE organizations
SET name = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteOrganization :exec
UPDATE organizations
SET deleted_at = now()
WHERE id = $1;

-- name: DeleteOrganizationByLogin :exec
UPDATE organizations
SET deleted_at = now()
WHERE login = $1 AND deleted_at IS NULL;
