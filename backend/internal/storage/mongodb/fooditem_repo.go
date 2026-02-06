package mongodb

import (
	"context"
	"errors"
	"time"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
	"github.com/qoofa/AI-Recommendation-System/internal/domain/food"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FoodRepository struct {
	collection *mongo.Collection
}

func NewFoodRepository(db *mongo.Database) *FoodRepository {
	return &FoodRepository{
		collection: db.Collection("food_items"),
	}
}

func (r *FoodRepository) Save(ctx context.Context, item *food.FoodItemModel) (string, error) {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now

	result, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return "", appErr.Wrap(appErr.Internal, "internal error", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", appErr.Wrap(appErr.Internal, "internal error", nil)
	}

	return oid.Hex(), nil
}

func (r *FoodRepository) FindByID(ctx context.Context, id string) (*food.FoodItemModel, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, appErr.New(appErr.BadRequest, "invalid id")
	}

	var m food.FoodItemModel
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&m)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, appErr.New(appErr.NotFound, "item not found")
		}
		return nil, err
	}

	return &m, nil
}

func (r *FoodRepository) FindByKeyword(ctx context.Context, query string) ([]food.FoodItemModel, error) {
	filter := bson.M{
		"$or": bson.A{
			bson.M{"name": bson.M{"$regex": query, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}
	defer cursor.Close(ctx)

	var foodItem []food.FoodItemModel
	if err := cursor.All(ctx, &foodItem); err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	return foodItem, nil
}

func (r *FoodRepository) FindBySemantic(ctx context.Context, embedding []float64) ([]food.FoodItemModel, error) {
	pipeline := bson.A{
		bson.M{
			"$vectorSearch": bson.M{
				"index":         "vector_index",
				"path":          "embedding",
				"queryVector":   embedding,
				"numCandidates": 50,
				"limit":         10,
			},
		},
		bson.M{
			"$project": bson.M{
				"name":        1,
				"description": 1,
				"image":       1,
				"category":    1,
				"score": bson.M{
					"$meta": "vectorSearchScore",
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}
	defer cursor.Close(ctx)

	var foodItem []food.FoodItemModel
	if err := cursor.All(ctx, &foodItem); err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	return foodItem, nil
}

func (r *FoodRepository) toDto(d *food.FoodItemModel) *FoodItemModel {
	if d == nil {
		return nil
	}

	dbModel := &FoodItemModel{
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Image:       d.Image,
		Category:    d.Category,
		SalesCount:  d.SalesCount,
		Embedding:   d.Embedding,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

	if d.Combos != nil {
		dbModel.Combos = make([]Combo, len(d.Combos))
		for i, v := range d.Combos {
			dbModel.Combos[i] = Combo(v)
		}
	}

	if d.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(d.ID); err == nil {
			dbModel.ID = oid
		}
	}

	return dbModel
}

func (r *FoodRepository) toDomain(d *FoodItemModel) *food.FoodItemModel {
	if d == nil {
		return nil
	}

	model := &food.FoodItemModel{
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Image:       d.Image,
		Category:    d.Category,
		SalesCount:  d.SalesCount,
		Embedding:   d.Embedding,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

	if d.Combos != nil {
		model.Combos = make([]food.Combo, len(d.Combos))
		for i, v := range d.Combos {
			model.Combos[i] = food.Combo(v)
		}
	}

	if !d.ID.IsZero() {
		model.ID = d.ID.Hex()
	}

	return model
}
