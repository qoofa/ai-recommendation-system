package mongodb

import (
	"context"
	"log"
	"strings"
	"time"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"

	orderembeddings "github.com/qoofa/AI-Recommendation-System/internal/domain/orderEmbeddings"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OrderEmbeddingRepository struct {
	collection *mongo.Collection
}

func NewOrderEmbeddingRepository(db *mongo.Database) *OrderEmbeddingRepository {
	if db == nil {
		return nil
	}

	repo := &OrderEmbeddingRepository{
		collection: db.Collection("order_embeddings"),
	}

	repo.ensureIndexes()

	return repo
}

func (r *OrderEmbeddingRepository) Save(ctx context.Context, d *orderembeddings.OrderEmbedding) (string, error) {
	if d == nil {
		return "", appErr.New(appErr.BadRequest, "invalid input")
	}

	now := time.Now()
	d.CreatedAt = now
	d.UpdatedAt = now
	payload := r.toDto(d)

	result, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return "", appErr.Wrap(appErr.Internal, "internal error", err)
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", appErr.New(appErr.Internal, "internal error")
	}

	return oid.Hex(), nil
}

func (r *OrderEmbeddingRepository) FindBySemantic(ctx context.Context, embedding []float64) ([]orderembeddings.OrderEmbedding, error) {
	pipeline := bson.A{
		bson.M{
			"$vectorSearch": bson.M{
				"index":         "order_vector_index",
				"path":          "embedding",
				"queryVector":   embedding,
				"numCandidates": 100,
				"limit":         20,
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}
	defer cursor.Close(ctx)

	var data []OrderEmbeddingModel
	if err := cursor.All(ctx, &data); err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	return r.toDomains(data), nil
}

func (r *OrderEmbeddingRepository) ensureIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = r.collection.Database().CreateCollection(ctx, r.collection.Name())

	cmd := bson.D{
		{Key: "createSearchIndexes", Value: "order_embeddings"},
		{Key: "indexes", Value: bson.A{
			bson.M{
				"name": "vector_index",
				"type": "vectorSearch",
				"definition": bson.M{
					"fields": bson.A{
						bson.M{
							"type":          "vector",
							"path":          "embedding",
							"numDimensions": 768,
							"similarity":    "cosine",
						},
					},
				},
			},
		}},
	}

	err := r.collection.Database().RunCommand(ctx, cmd).Err()
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Printf("DATABASE ERROR: Warning: vector index setup failed: %v", err)
	}
}

func (r *OrderEmbeddingRepository) toDto(d *orderembeddings.OrderEmbedding) *OrderEmbeddingModel {
	if d == nil {
		return nil
	}

	dbModel := &OrderEmbeddingModel{
		Embedding: d.Embedding,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	if d.ID != "" {
		if oid, err := bson.ObjectIDFromHex(d.ID); err == nil {
			dbModel.ID = oid
		}
	}

	if len(d.Items) > 0 {
		j := 0
		dbModel.Items = make([]bson.ObjectID, len(d.Items))
		for i := range d.Items {
			if oid, err := bson.ObjectIDFromHex(d.Items[i]); err == nil {
				dbModel.Items[i] = oid
				j++
			}
		}
	}

	return dbModel
}

func (r *OrderEmbeddingRepository) toDomain(d *OrderEmbeddingModel) *orderembeddings.OrderEmbedding {
	if d == nil {
		return nil
	}

	model := &orderembeddings.OrderEmbedding{
		Embedding: d.Embedding,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	if !d.ID.IsZero() {
		model.ID = d.ID.Hex()
	}

	if len(d.Items) > 0 {
		j := 0
		model.Items = make([]string, len(d.Items))
		for i := range d.Items {
			if !d.Items[i].IsZero() {
				model.Items[j] = d.Items[i].Hex()
				j++
			}
		}
	}

	return model
}

func (r *OrderEmbeddingRepository) toDomains(models []OrderEmbeddingModel) []orderembeddings.OrderEmbedding {
	result := make([]orderembeddings.OrderEmbedding, len(models))
	for i := range models {
		result[i] = *r.toDomain(&models[i])
	}
	return result
}
