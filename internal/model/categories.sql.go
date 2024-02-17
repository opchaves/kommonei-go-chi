// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: categories.sql

package model

import (
	"context"

	"github.com/google/uuid"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (
  name, user_id, workspace_id, cat_type, description, icon, color
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
`

type CreateCategoryParams struct {
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"user_id"`
	WorkspaceID uuid.UUID `json:"workspace_id"`
	CatType     string    `json:"cat_type"`
	Description *string   `json:"description"`
	Icon        *string   `json:"icon"`
	Color       *string   `json:"color"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.Exec(ctx, createCategory,
		arg.Name,
		arg.UserID,
		arg.WorkspaceID,
		arg.CatType,
		arg.Description,
		arg.Icon,
		arg.Color,
	)
	return err
}

const deleteCategoriesByWorkspace = `-- name: DeleteCategoriesByWorkspace :exec
DELETE FROM categories WHERE workspace_id = $1
`

func (q *Queries) DeleteCategoriesByWorkspace(ctx context.Context, workspaceID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCategoriesByWorkspace, workspaceID)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const getCategoriesByUser = `-- name: GetCategoriesByUser :many
SELECT id, name, description, icon, color, cat_type, user_id, workspace_id, created_at, updated_at FROM categories WHERE user_id = $1
`

func (q *Queries) GetCategoriesByUser(ctx context.Context, userID uuid.UUID) ([]*Category, error) {
	rows, err := q.db.Query(ctx, getCategoriesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Icon,
			&i.Color,
			&i.CatType,
			&i.UserID,
			&i.WorkspaceID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoriesByWorkspace = `-- name: GetCategoriesByWorkspace :many
SELECT id, name, description, icon, color, cat_type, user_id, workspace_id, created_at, updated_at FROM categories WHERE workspace_id = $1
`

func (q *Queries) GetCategoriesByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*Category, error) {
	rows, err := q.db.Query(ctx, getCategoriesByWorkspace, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Icon,
			&i.Color,
			&i.CatType,
			&i.UserID,
			&i.WorkspaceID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id, name, description, icon, color, cat_type, user_id, workspace_id, created_at, updated_at FROM categories WHERE id = $1
`

func (q *Queries) GetCategoryByID(ctx context.Context, id uuid.UUID) (*Category, error) {
	row := q.db.QueryRow(ctx, getCategoryByID, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Icon,
		&i.Color,
		&i.CatType,
		&i.UserID,
		&i.WorkspaceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories SET
  name = $1,
  description = $2,
  icon = $3,
  color = $4
WHERE id = $5
`

type UpdateCategoryParams struct {
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Icon        *string   `json:"icon"`
	Color       *string   `json:"color"`
	ID          uuid.UUID `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.Exec(ctx, updateCategory,
		arg.Name,
		arg.Description,
		arg.Icon,
		arg.Color,
		arg.ID,
	)
	return err
}
