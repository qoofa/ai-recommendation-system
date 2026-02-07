package food

import (
	"context"
	"time"
)

type Combo struct {
	ItemID string
	Count  int
}

type FoodItemModel struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Image       string
	Category    string
	SalesCount  int

	Embedding []float64
	Score     float64

	Combos []Combo

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Search(ctx context.Context, query string) ([]FoodItemModel, error)
	Create(ctx context.Context, d FoodItemModel) (string, error)
}

type Repository interface {
	Save(ctx context.Context, item *FoodItemModel) (string, error)
	InsertMany(ctx context.Context, item []FoodItemModel) ([]string, error)
	FindByID(ctx context.Context, id string) (*FoodItemModel, error)
	FindByKeyword(ctx context.Context, query string) ([]FoodItemModel, error)
	FindBySemantic(ctx context.Context, embedding []float64) ([]FoodItemModel, error)
}
