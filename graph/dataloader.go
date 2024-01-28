package graph

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/lil-shimon/lil-gqland/graph/model"
	"github.com/lil-shimon/lil-gqland/graph/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(Srv services.Services) *Loaders {
	userBatcher := &userBatcher{Srv: Srv}
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.UserService
}

func (u *userBatcher) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	// 結果の配列を作成
	// ここでは、IDの数だけ結果の配列を作成している
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("User not found"),
		}
	}

	// IDをキーにして、インデックスを作成
	indexes := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexes[ID] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)

	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexes[user.ID]] = rsl
	}

	return results
}
