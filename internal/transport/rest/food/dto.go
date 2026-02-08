package food

import (
	"time"

	"github.com/qoofa/AI-Recommendation-System/internal/core"
)

type createDto struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"required,max=500"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	SalesCount  int     `json:"sales_count" validate:"gte=0"`
}

func (d *createDto) toDomain() *core.FoodItemModel {
	return &core.FoodItemModel{
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Image:       d.Image,
		Category:    d.Category,
		SalesCount:  d.SalesCount,
	}
}

type SearchResponseDto struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Image       string    `json:"image,omitempty"`
	Category    string    `json:"category,omitempty"`
	SalesCount  int       `json:"sales_count,omitempty"`
	Score       float64   `json:"score,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func TOSearchResponse(m core.FoodItemModel) SearchResponseDto {
	return SearchResponseDto{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Image:       m.Image,
		Category:    m.Category,
		SalesCount:  m.SalesCount,
		Score:       m.Score,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
