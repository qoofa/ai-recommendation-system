package main

import (
	"log"
	"net/http"

	"github.com/qoofa/AI-Recommendation-System/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env: ", err)
	}

	router, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
