package main

import (
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller"
	"movieexample.com/metadata/internal/handler"
	"movieexample.com/metadata/internal/repository"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := repository.New()
	ctrl := controller.New(repo)
	h := handler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicf("Failed to start server:n%v\n", err)
	}

}
