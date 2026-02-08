package mongodb

import (
	"context"
	"log"
	"strings"
	"time"

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

	repo.ensureOrderEmbeddingIndexes()

	return repo
}

func (r *OrderEmbeddingRepository) ensureOrderEmbeddingIndexes() {
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
