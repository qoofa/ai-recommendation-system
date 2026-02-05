package app

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/storage/mongodb"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/food"
)

type App struct {
	Router *chi.Mux
}

func New() (*chi.Mux, error) {
	_, err := mongodb.New(os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	foodHandler := food.NewFoodHandler()

	return rest.NewRouter(foodHandler), nil
}
