package food

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Combo struct {
	ItemID string
	Count  int
}

type FoodItemModel struct {
	ID          bson.ObjectID
	Name        string
	Description string
	Price       float64
	Image       string
	Category    string
	SalesCount  int

	Embedding []float64

	Combos []Combo

	CreatedAt time.Time
	UpdatedAt time.Time
}
