package services

import (
	"context"

	"github.com/lil-shimon/lil-gqland/graph/db"
	"github.com/lil-shimon/lil-gqland/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

// convertUser converts a db.User(sqlboiler) to a model.User(graphql)
func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	return convertUser(user), nil
}
