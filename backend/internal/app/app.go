package app

import (
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/qoofa/AI-Recommendation-System/internal/infrastructure/embeddings"
	"github.com/qoofa/AI-Recommendation-System/internal/storage/mongodb"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest"
	foodH "github.com/qoofa/AI-Recommendation-System/internal/transport/rest/food"

	"github.com/qoofa/AI-Recommendation-System/internal/domain/food"
)

type App struct {
	Router *chi.Mux
}

func New() (*chi.Mux, error) {
	db, err := mongodb.New(os.Getenv("DB_DSN"), os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "DATABASE error: %v", err)
	}

	embedder := embeddings.NewPythonProvider(os.Getenv("EMBEDDING_SERVER_URL"))

	foodRepo := mongodb.NewFoodRepository(db)
	foodServie := food.New(foodRepo, embedder)
	foodHandler := foodH.NewFoodHandler(foodServie)

	return rest.NewRouter(foodHandler), nil
}
