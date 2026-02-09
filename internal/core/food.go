package core

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

type FoodService interface {
	Find(ctx context.Context) ([]FoodItemModel, error)
	Search(ctx context.Context, query string) ([]FoodItemModel, error)
	Create(ctx context.Context, d FoodItemModel) (string, error)
	Recommend(ctx context.Context, itemId string) ([]FoodItemModel, error)
}
