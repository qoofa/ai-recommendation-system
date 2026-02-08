package orderembeddings

import "time"

type OrderEmbedding struct {
	ID        string
	Embedding []float64
	Items     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
