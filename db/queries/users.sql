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
