package main

import (
	"log"
	"net/http"

	"github.com/gkatanacio/card-deck-api/internal/deck"
	"github.com/gkatanacio/card-deck-api/internal/handler"
	"github.com/gkatanacio/card-deck-api/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("server startup")

	db, err := storage.NewDb(storage.NewConfig())
	if err != nil {
		panic(err)
	}

	service := deck.NewService(deck.NewDatabase(db))

	r := mux.NewRouter()
	r.HandleFunc("/decks", handler.NewCreateDeckHandler(service).Handle).Methods(http.MethodPost)
	r.HandleFunc("/decks/{id}", handler.NewGetDeckHandler(service).Handle).Methods(http.MethodGet)
	r.HandleFunc("/decks/{id}/cards", handler.NewDeleteCardsHandler(service).Handle).Methods(http.MethodDelete)

	port := ":8080"
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
