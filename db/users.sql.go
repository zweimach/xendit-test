// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package db

import (
	"context"
)

const listUsersByOrganizationLogin = `-- name: ListUsersByOrganizationLogin :many
SELECT u.id, u.login, u.name, u.followers, u.follows, u.avatar_url, u.created_at, u.updated_at, u.deleted_at FROM users AS u
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
OFFSET $3
`

type ListUsersByOrganizationLoginParams struct {
	Login  string
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsersByOrganizationLogin(ctx context.Context, arg ListUsersByOrganizationLoginParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByOrganizationLogin, arg.Login, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Name,
			&i.Followers,
			&i.Follows,
			&i.AvatarUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
