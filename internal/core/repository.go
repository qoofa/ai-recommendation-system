package core

import "context"

type FoodRepository interface {
	Save(ctx context.Context, item *FoodItemModel) (string, error)
	Find(ctx context.Context) ([]FoodItemModel, error)
	InsertMany(ctx context.Context, item []FoodItemModel) ([]string, error)
	FindByIds(ctx context.Context, ids []string) ([]FoodItemModel, error)
	FindByID(ctx context.Context, id string) (*FoodItemModel, error)
	FindByKeyword(ctx context.Context, query string) ([]FoodItemModel, error)
	FindBySemantic(ctx context.Context, embedding []float64) ([]FoodItemModel, error)
}

type OrderEmbeddingRepository interface {
	Save(ctx context.Context, d *OrderEmbedding) (string, error)
	FindBySemantic(ctx context.Context, embedding []float64) ([]OrderEmbedding, error)
}
