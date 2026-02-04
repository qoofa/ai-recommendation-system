package mongodb

import (
	"context"
	"errors"
	"time"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Combo struct {
	ItemID string `bson:"itemId"`
	Count  int    `bson:"count"`
}

type FoodItemModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Price       float64       `bson:"price"`
	Image       string        `bson:"image"`
	Category    string        `bson:"category"`
	SalesCount  int           `bson:"salesCount"`

	Embedding []float64 `bson:"embedding"`

	Combos []Combo `bson:"combos"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type FoodRepository struct {
	collection *mongo.Collection
}

func NewFoodRepository(db *mongo.Database) *FoodRepository {
	return &FoodRepository{
		collection: db.Collection("food_items"),
	}
}

func (r *FoodRepository) Save(ctx context.Context, item FoodItemModel) error {
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, item)
	return err
}

func (r *FoodRepository) FindByID(ctx context.Context, id string) (*FoodItemModel, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, appErr.New(appErr.BadRequest, "invalid id")
	}

	var m FoodItemModel
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&m)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, appErr.New(appErr.NotFound, "item not found")
		}
		return nil, err
	}

	return &m, nil
}

func (r *FoodRepository) FindByRejex(ctx context.Context, query string) (*mongo.Cursor, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"name": bson.M{"$regex": query, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}
	return r.collection.Find(ctx, filter)
}
