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

type Service interface {
	Train(ctx context.Context, items []string) (string, error)
}

type Repository interface {
	Save(ctx context.Context, d *OrderEmbedding) (string, error)
	FindBySemantic(ctx context.Context, embedding []float64) ([]OrderEmbedding, error)
}
