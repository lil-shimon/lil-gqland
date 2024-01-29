package services

import (
	"context"

	"github.com/lil-shimon/lil-gqland/graph/db"
	"github.com/lil-shimon/lil-gqland/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type brandService struct {
	exec boil.ContextExecutor
}

func convertBrand(brand *db.Brand) *model.Brand {
	return &model.Brand{
		ID:   brand.ID,
		Name: brand.Name,
	}
}

func (b *brandService) GetBrandById(ctx context.Context, id string) (*model.Brand, error) {
	brand, err := db.FindBrand(ctx, b.exec, id)
	if err != nil {
		return nil, err
	}

	return convertBrand(brand), nil
}
