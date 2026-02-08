package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/food"
	orderembedding "github.com/qoofa/AI-Recommendation-System/internal/transport/rest/orderEmbedding"
)

func NewRouter(foodHandler *food.FoodHandler, OrderEmbeddingHandler *orderembedding.OrderEmbeddingHanlder) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/food", foodHandler.Find)
		r.Get("/food/search", foodHandler.Search)
		r.Post("/food", foodHandler.Create)
		r.Post("/order/train", OrderEmbeddingHandler.Train)
	})

	return r
}
