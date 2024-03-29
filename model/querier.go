// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package model

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (*Account, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) error
	CreateToken(ctx context.Context, arg CreateTokenParams) (*Token, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (*Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (*Workspace, error)
	CreateWorkspaceUser(ctx context.Context, arg CreateWorkspaceUserParams) (*WorkspacesUser, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	DeleteCategoriesByWorkspace(ctx context.Context, workspaceID uuid.UUID) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	DeleteTokenByID(ctx context.Context, id int32) error
	DeleteTransaction(ctx context.Context, arg DeleteTransactionParams) error
	DeleteWorkspace(ctx context.Context, arg DeleteWorkspaceParams) error
	GetAccountByID(ctx context.Context, id uuid.UUID) (*Account, error)
	GetAccountsByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*Account, error)
	GetCategoriesByUser(ctx context.Context, userID uuid.UUID) ([]*Category, error)
	GetCategoriesByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*Category, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*Category, error)
	GetDefaultUserWorkspace(ctx context.Context, id uuid.UUID) (*GetDefaultUserWorkspaceRow, error)
	GetToken(ctx context.Context, token string) (*Token, error)
	GetTokenById(ctx context.Context, id int32) (*Token, error)
	GetTokensByUser(ctx context.Context, userID uuid.UUID) ([]*Token, error)
	GetTransactionById(ctx context.Context, arg GetTransactionByIdParams) (*Transaction, error)
	GetTransactionsByUser(ctx context.Context, userID uuid.UUID) ([]*Transaction, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
	GetWorkspacesByUser(ctx context.Context, arg GetWorkspacesByUserParams) ([]*Workspace, error)
	IsEmailTaken(ctx context.Context, email string) (int32, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error
	UpdateToken(ctx context.Context, arg UpdateTokenParams) error
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (*Transaction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error)
	UpdateWorkspace(ctx context.Context, arg UpdateWorkspaceParams) (*Workspace, error)
}

var _ Querier = (*Queries)(nil)
