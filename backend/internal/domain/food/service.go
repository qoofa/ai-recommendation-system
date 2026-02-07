package food

import (
	"context"
	"fmt"
	"sort"

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

func (s *service) Search(ctx context.Context, query string) ([]FoodItemModel, error) {
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

	type combined struct {
		item  FoodItemModel
		score float64
	}

	combinedMap := make(map[string]combined, 0)
	for i := range keywordResult {
		id := keywordResult[i].ID
		combinedMap[id] = combined{
			item:  keywordResult[i],
			score: 0.3,
		}
	}

	for i := range semanticResults {
		id := semanticResults[i].ID
		semanticScore := semanticResults[i].Score

		if item, ok := combinedMap[id]; ok {
			item.score += 0.5 * semanticScore
		} else {
			combinedMap[id] = combined{
				item:  semanticResults[i],
				score: 0.7 * semanticScore,
			}
		}
	}

	combinedSlice := make([]combined, len(combinedMap))
	for _, v := range combinedMap {
		combinedSlice = append(combinedSlice, v)
	}

	sort.Slice(combinedSlice, func(i, j int) bool {
		return combinedSlice[i].score > combinedSlice[j].score
	})

	limit := 3
	if len(combinedSlice) < limit {
		limit = len(combinedSlice)
	}

	result := make([]FoodItemModel, limit)
	for i := 0; i < limit; i++ {
		combinedSlice[i].item.Embedding = nil
		result[i] = combinedSlice[i].item
	}

	return result, nil
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
