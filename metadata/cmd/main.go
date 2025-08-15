package main

import (
	"log"
	"net/http"

	"github.com/stevan1008/movieGolang/metadata/internal/controller/metadata"
	httphandler "github.com/stevan1008/movieGolang/metadata/internal/handler/http"
	"github.com/stevan1008/movieGolang/metadata/internal/repository/memory"
)

func main() {
	log.Printf("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.HandleFunc("/metadata", h.GetMetadata)
	if err := http.ListenAndServe(":8072", nil); err != nil {
		panic(err)
	}
}
