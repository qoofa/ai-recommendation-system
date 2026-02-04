package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/food"
)

func NewRouter(foodHandler *food.FoodHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/search", foodHandler.Search)
	})

	return r
}
