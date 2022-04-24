-- name: CreateMembership :one
INSERT INTO memberships (
    user_id,
    organization_id
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetMembershipByID :one
SELECT * FROM memberships
WHERE user_id = $1 AND organization_id = $2 AND deleted_at IS NULL
LIMIT 1;

-- name: ListMemberships :many
SELECT * FROM memberships
WHERE deleted_at IS NULL
LIMIT $1
OFFSET $2;

-- name: DeleteMembership :exec
UPDATE comments
SET deleted_at = now()
WHERE user_id = $1 AND organization_id = $2;
