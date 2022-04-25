-- name: CreateCommentByOrganizationLogin :one
INSERT INTO comments (
    organization_id,
    text
)
SELECT o.id, $1 FROM organizations AS o
WHERE o.login = $2 AND deleted_at IS NULL
LIMIT 1
RETURNING *;

-- name: ListCommentsByOrganizationLogin :many
SELECT c.* FROM comments AS c
INNER JOIN organizations AS o
ON o.id = c.organization_id
WHERE o.login = $1 AND o.deleted_at IS NULL AND c.deleted_at IS NULL
ORDER BY c.id
LIMIT $2
OFFSET $3;

-- name: DeleteCommentsByOrganizationLogin :exec
UPDATE comments AS c
SET deleted_at = now()
FROM organizations AS o
WHERE c.organization_id = o.id
    AND o.login = $1
    AND o.deleted_at IS NULL
    AND c.deleted_at IS NULL;
