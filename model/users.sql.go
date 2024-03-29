// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package model

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  password,
  verified,
  verification_token,
  avatar
) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, first_name, last_name, email, password, verified, verification_token, avatar, created_at, updated_at
`

type CreateUserParams struct {
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Email             string  `json:"email"`
	Password          string  `json:"password"`
	Verified          bool    `json:"verified"`
	VerificationToken *string `json:"verification_token"`
	Avatar            string  `json:"avatar"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Verified,
		arg.VerificationToken,
		arg.Avatar,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Verified,
		&i.VerificationToken,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, verified, verification_token, avatar, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Verified,
		&i.VerificationToken,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, first_name, last_name, email, password, verified, verification_token, avatar, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (*User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Verified,
		&i.VerificationToken,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name, last_name, email, password, verified, verification_token, avatar, created_at, updated_at FROM users ORDER BY id
`

func (q *Queries) GetUsers(ctx context.Context) ([]*User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.Verified,
			&i.VerificationToken,
			&i.Avatar,
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

const isEmailTaken = `-- name: IsEmailTaken :one
SELECT 1 FROM users WHERE email = $1
`

func (q *Queries) IsEmailTaken(ctx context.Context, email string) (int32, error) {
	row := q.db.QueryRow(ctx, isEmailTaken, email)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET
  first_name = coalesce($1, first_name),
  last_name = coalesce($2, last_name),
  email = coalesce($3, email),
  password = coalesce($4, password),
  verified = coalesce($5, verified),
  verification_token = coalesce($6, verification_token),
  updated_at = now()
WHERE id = $7
RETURNING id, first_name, last_name, email, password, verified, verification_token, avatar, created_at, updated_at
`

type UpdateUserParams struct {
	FirstName         *string   `json:"first_name"`
	LastName          *string   `json:"last_name"`
	Email             *string   `json:"email"`
	Password          *string   `json:"password"`
	Verified          *bool     `json:"verified"`
	VerificationToken *string   `json:"verification_token"`
	ID                uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Verified,
		arg.VerificationToken,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Verified,
		&i.VerificationToken,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
