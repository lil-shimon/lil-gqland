package services

import (
	"fmt"

	"github.com/lil-shimon/lil-gqland/graph/db"
	"github.com/lil-shimon/lil-gqland/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/net/context"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

// convertRespository converts a db.Repository(sqlboiler) to a model.Repository(graphql)
func convertRespository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:    repo.ID,
		Name:  repo.Name,
		Owner: &model.User{ID: repo.Owner},
	}
}

func (r *repositoryService) GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)

	fmt.Println("name", name, "owner", owner)
	fmt.Println("repo", repo, "err", err)

	if err != nil {
		return nil, err
	}

	return convertRespository(repo), nil
}
