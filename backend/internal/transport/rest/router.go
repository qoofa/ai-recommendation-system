package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/food"
)

func NewRouter(foodHandler *food.FoodHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/food/search", foodHandler.Search)
		r.Post("/food", foodHandler.Create)
	})

	return r
}
