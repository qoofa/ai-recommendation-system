package service

import (
	"context"

	"github.com/qoofa/AI-Recommendation-System/internal/core"
	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type orderEmbeddingService struct {
	repo     core.OrderEmbeddingRepository
	foodRepo core.FoodRepository
}

func NewOrderEmbeddingService(r core.OrderEmbeddingRepository, f core.FoodRepository) *orderEmbeddingService {
	return &orderEmbeddingService{
		repo:     r,
		foodRepo: f,
	}
}

func (s *orderEmbeddingService) Train(ctx context.Context, items []string) (string, error) {
	if len(items) < 1 {
		return "", appErr.New(appErr.BadRequest, "invalid items")
	}

	foodItems, err := s.foodRepo.FindByIds(ctx, items)
	if err != nil {
		return "", err
	}

	if len(foodItems) == 0 {
		return "", appErr.New(appErr.NotFound, "not items found")
	}

	var itemEmbedding [][]float64

	for i := range foodItems {
		itemEmbedding = append(itemEmbedding, foodItems[i].Embedding)
	}

	dims := len(itemEmbedding[0])

	orderEmbedding := make([]float64, dims)

	for _, i := range itemEmbedding {
		for idx, v := range i {
			orderEmbedding[idx] += v
		}
	}

	for i := range dims {
		orderEmbedding[i] /= float64(len(itemEmbedding))
	}

	payload := &core.OrderEmbedding{
		Items:     items,
		Embedding: orderEmbedding,
	}

	saved, err := s.repo.Save(ctx, payload)
	if err != nil {
		return "", err
	}

	return saved, nil
}
