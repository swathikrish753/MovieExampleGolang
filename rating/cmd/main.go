package main

import (
	"log"
	"net/http"

	"movieexample.com/rating/internal/controller/rating"
	handler "movieexample.com/rating/internal/handler/http"
	repository "movieexample.com/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service...")
	repo := repository.New()
	ctrl := rating.New(repo)
	h := handler.NewHandler(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
