package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	//env := godotenv.Load()
	godotenv.Load(".env")

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

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(portString)

}
