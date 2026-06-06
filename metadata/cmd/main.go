package main

import (
	"log"
	"net/http"

	metadata "movieexample.com/metadata/internal/controller"
	httphandler "movieexample.com/metadata/internal/handler"

	// removed unused http handler import
	memory "movieexample.com/metadata/internal/repository"
)

func main() {

	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.HandleFunc("/metadata", h.GetMetadata)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}
