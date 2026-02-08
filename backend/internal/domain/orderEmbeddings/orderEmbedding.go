package orderembeddings

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

type Repository interface {
	Save(ctx context.Context, d *OrderEmbedding) (string, error)
}
