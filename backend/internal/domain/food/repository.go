package food

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, item FoodItemModel) (string, error)
	FindByID(ctx context.Context, id string) (*FoodItemModel, error)
	FindByKeyword(ctx context.Context, query string) ([]FoodItemModel, error)
	FindBySemantic(ctx context.Context, embedding []float64) ([]FoodItemModel, error)
}
