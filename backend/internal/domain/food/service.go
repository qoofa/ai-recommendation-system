package food

import (
	"context"
	"fmt"

	"github.com/qoofa/AI-Recommendation-System/internal/domain/embedding"
)

type service struct {
	repo      Repository
	embedding embedding.Embedder
}

func NewService(r Repository, e embedding.Embedder) *service {
	return &service{
		repo:      r,
		embedding: e,
	}
}

func (s *service) Search(ctx context.Context, query string) (*[]FoodItemModel, error) {
	if query == "" {
		return nil, nil
	}

	embedding, err := s.embedding.GetEmbedding(ctx, query)
	if err != nil {
		return nil, err
	}

	keywordResult, err := s.repo.FindByKeyword(ctx, query)
	if err != nil {
		return nil, err
	}

	semanticResults, err := s.repo.FindBySemantic(ctx, embedding)
	if err != nil {
		return nil, err
	}

	fmt.Println(embedding)
	fmt.Println(keywordResult)
	fmt.Println(semanticResults)

	return nil, nil
}

func (s *service) Create(ctx context.Context, d FoodItemModel) (string, error) {
	query := fmt.Sprintf("%s. %s", d.Name, d.Description)

	embedding, err := s.embedding.GetEmbedding(ctx, query)
	if err != nil {
		return "", err
	}

	d.Embedding = embedding

	return s.repo.Save(ctx, &d)
}
