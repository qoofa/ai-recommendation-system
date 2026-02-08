package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/qoofa/AI-Recommendation-System/internal/core"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type foodService struct {
	repo      core.FoodRepository
	orderRepo core.OrderEmbeddingRepository
	embedding core.Embedder
}

func NewFoodService(r core.FoodRepository, o core.OrderEmbeddingRepository, e core.Embedder) *foodService {
	return &foodService{
		repo:      r,
		orderRepo: o,
		embedding: e,
	}
}

func (s *foodService) Find(ctx context.Context) ([]core.FoodItemModel, error) {
	result, err := s.repo.Find(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *foodService) Search(ctx context.Context, query string) ([]core.FoodItemModel, error) {
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
		item  core.FoodItemModel
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

	result := make([]core.FoodItemModel, limit)
	for i := 0; i < limit; i++ {
		combinedSlice[i].item.Embedding = nil
		result[i] = combinedSlice[i].item
	}

	return result, nil
}

func (s *foodService) Create(ctx context.Context, d core.FoodItemModel) (string, error) {
	query := fmt.Sprintf("%s. %s", d.Name, d.Description)

	embedding, err := s.embedding.GetEmbedding(ctx, query)
	if err != nil {
		return "", err
	}

	d.Embedding = embedding

	return s.repo.Save(ctx, &d)
}

func (s *foodService) Recommend(ctx context.Context, itemId string) ([]core.FoodItemModel, error) {
	if itemId == "" {
		return nil, appErr.New(appErr.BadRequest, "invalid id")
	}

	item, err := s.repo.FindByID(ctx, itemId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, appErr.New(appErr.BadRequest, "item not found")
	}

	matches, err := s.orderRepo.FindBySemantic(ctx, item.Embedding)
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)

	for i := range matches {
		for _, id := range matches[i].Items {
			if id != itemId {
				counts[id]++
			}
		}
	}

	type entry struct {
		Key   string
		Value int
	}

	var sortedEntries []entry
	for k, v := range counts {
		sortedEntries = append(sortedEntries, entry{Key: k, Value: v})
	}

	sort.Slice(sortedEntries, func(i, j int) bool {
		return sortedEntries[i].Value > sortedEntries[j].Value
	})

	limit := 6
	if len(sortedEntries) < 6 {
		limit = len(sortedEntries)
	}

	recommentedIds := make([]string, limit)

	for i := range limit {
		recommentedIds[i] = sortedEntries[i].Key
	}

	recommented, err := s.repo.FindByIds(ctx, recommentedIds)
	if err != nil {
		return nil, err
	}

	return recommented, nil
}
