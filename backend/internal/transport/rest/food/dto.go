package food

import "github.com/qoofa/AI-Recommendation-System/internal/domain/food"

type createDto struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"required,max=500"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	SalesCount  int     `json:"sales_count" validate:"gte=0"`
}

func (d *createDto) toDomain() *food.FoodItemModel {
	return &food.FoodItemModel{
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Image:       d.Image,
		Category:    d.Category,
		SalesCount:  d.SalesCount,
	}
}
