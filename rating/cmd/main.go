package main

import (
	"log"
	"net/http"

	rating "movieexample.com/rating/internal/controller"
	httphandler "movieexample.com/rating/internal/handler/http"
	memory "movieexample.com/rating/internal/repository"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctr := rating.Controller{}
	ctrl := ctr.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

}
