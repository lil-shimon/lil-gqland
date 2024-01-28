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

// convertUserSlice converts a db.UserSlice(sqlboiler) to a model.UserSlice(graphql)
func convertUserSlice(users db.UserSlice) []*model.User {
	result := make([]*model.User, 0, len(users))
	for _, user := range users {
		result = append(result, convertUser(user))
	}

	return result
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

// GetUserByID returns a user by its id
func (u *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := db.FindUser(ctx, u.exec, id, db.UserTableColumns.ID, db.UserTableColumns.Name)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func (u *userService) ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error) {
	users, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.IN(IDs),
	).All(ctx, u.exec)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	return convertUserSlice(users), nil
}
