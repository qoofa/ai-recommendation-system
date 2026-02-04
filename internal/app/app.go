package app

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/storage/mongodb"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest"
)

type App struct {
	Router *chi.Mux
}

func New() (*chi.Mux, error) {
	_, err := mongodb.New(os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	return rest.NewRouter(), nil
}
