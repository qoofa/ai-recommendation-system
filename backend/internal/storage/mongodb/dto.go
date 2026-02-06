package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Combo struct {
	ItemID string `bson:"itemId,omitempty"`
	Count  int    `bson:"count,omitempty"`
}

type FoodItemModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name,omitempty"`
	Description string        `bson:"description,omitempty"`
	Price       float64       `bson:"price,omitempty"`
	Image       string        `bson:"image,omitempty"`
	Category    string        `bson:"category,omitempty"`
	SalesCount  int           `bson:"salesCount,omitempty"`

	Embedding []float64 `bson:"embedding,omitempty"`

	Combos []Combo `bson:"combos,omitempty"`

	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
