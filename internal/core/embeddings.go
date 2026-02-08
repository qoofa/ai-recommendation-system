package core

import "context"

type Embedder interface {
	GetEmbedding(ctx context.Context, text string) ([]float64, error)
}
