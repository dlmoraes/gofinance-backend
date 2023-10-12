// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    title,
    type,
    description
) VALUES (
   $1, $2, $3, $4
) RETURNING id, user_id, title, type, description, created_at
`

type CreateCategoryParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.UserID,
		arg.Title,
		arg.Type,
		arg.Description,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT id, user_id, title, type, description, created_at FROM categories WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4
`

type GetCategoriesParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
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

const getCategory = `-- name: GetCategory :one
SELECT id, user_id, title, type, description, created_at FROM categories WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories SET title = $2, description = $3 WHERE id = $1 RETURNING id, user_id, title, type, description, created_at
`

type UpdateCategoryParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.Title, arg.Description)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
