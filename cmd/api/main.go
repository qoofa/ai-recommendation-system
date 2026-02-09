package main

import (
	"log"
	"net/http"
	"os"

	"github.com/qoofa/AI-Recommendation-System/internal/app"

	"github.com/joho/godotenv"
)

//	@title			AI Recommendation System API
//	@version		1.0
//	@description	This is a AI Recommendation System server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env: ", err)
	}

	router, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	log.Println("listening on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
