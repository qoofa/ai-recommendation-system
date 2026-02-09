package rest

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/qoofa/AI-Recommendation-System/docs"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/food"
	orderembedding "github.com/qoofa/AI-Recommendation-System/internal/transport/rest/orderEmbedding"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(foodHandler *food.FoodHandler, OrderEmbeddingHandler *orderembedding.OrderEmbeddingHanlder) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/food", foodHandler.Find)
		r.Post("/food", foodHandler.Create)
		r.Get("/food/search", foodHandler.Search)
		r.Get("/food/recommend", foodHandler.Recommend)	

		r.Post("/order/train", OrderEmbeddingHandler.Train)
	})

	return r
}
