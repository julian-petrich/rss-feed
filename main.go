package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"rss-feed/internal/database"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	//env := godotenv.Load()
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the .env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Couldn't connect to DB", err)
	}

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the .env")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUsers)

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(portString)

}
