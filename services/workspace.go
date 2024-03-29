package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opchaves/go-kom/model"
	"github.com/opchaves/go-kom/pkg/util"
)

type (
	WorkspaceService interface {
		Create(ctx context.Context, input *model.CreateWorkspaceParams) (*model.Workspace, error)
	}

	workspaceService struct {
		DB *pgxpool.Pool
		Q  *model.Queries
	}
)

func (s *workspaceService) Create(ctx context.Context, input *model.CreateWorkspaceParams) (*model.Workspace, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}

	qTx := s.Q.WithTx(tx)
	defer tx.Rollback(ctx)

	workspace, err := qTx.CreateWorkspace(ctx, *input)
	if err != nil {
		return nil, err
	}

	err = createCategories(ctx, qTx, workspace)
	if err != nil {
		return nil, err
	}

	err = createAccounts(ctx, qTx, workspace)
	if err != nil {
		return nil, err
	}

	_ = tx.Commit(ctx)

	return workspace, nil
}

func createCategories(ctx context.Context, db *model.Queries, workspace *model.Workspace) error {
	names := []model.CreateCategoryParams{
		{Name: "food", CatType: "expense"},
		{Name: "salary", CatType: "incom"},
		{Name: "health", CatType: "expense"},
		{Name: "transport", CatType: "expense"},
	}

	for _, c := range names {
		c.WorkspaceID = workspace.ID
		c.UserID = workspace.UserID
		if err := db.CreateCategory(ctx, c); err != nil {
			return err
		}
	}

	return nil
}

func createAccounts(ctx context.Context, db *model.Queries, workspace *model.Workspace) error {
	names := []model.CreateAccountParams{
		{Name: "checking", Balance: *util.ZERO},
		{Name: "savings", Balance: *util.ZERO},
	}

	for _, a := range names {
		a.WorkspaceID = workspace.ID
		a.UserID = workspace.UserID
		if _, err := db.CreateAccount(ctx, a); err != nil {
			return err
		}
	}

	return nil
}
