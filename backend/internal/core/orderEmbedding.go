package core

import (
	"context"
	"time"
)

type OrderEmbedding struct {
	ID        string
	Embedding []float64
	Items     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderEmbeddingService interface {
	Train(ctx context.Context, items []string) (string, error)
}
