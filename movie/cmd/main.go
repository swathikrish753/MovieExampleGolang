package main

import (
	"log"
	"net/http"

	"movieexample.com/movie/internal/controller/movie"
	metadataGateway "movieexample.com/movie/internal/gateway/metadata/http"
	ratingGateway "movieexample.com/movie/internal/gateway/rating/http"
	httphandler "movieexample.com/movie/internal/handler/http"
)

func main() {
	log.Println("Strating the movie service")
	metadataGateway := metadataGateway.New("localhost:8081")
	ratingGateway := ratingGateway.NewRgate("localhost:8082")
	ctrl := movie.NewController(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
