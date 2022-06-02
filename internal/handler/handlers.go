package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gkatanacio/card-deck-api/internal/deck"
	"github.com/gkatanacio/card-deck-api/internal/errs"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateDeckHandler handles: POST /decks?shuffle={bool}&cards={comma-separated card codes}.
type CreateDeckHandler struct {
	deckService deck.Servicer
}

func NewCreateDeckHandler(deckService deck.Servicer) *CreateDeckHandler {
	return &CreateDeckHandler{deckService}
}

func (h *CreateDeckHandler) Handle(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	shuffle := false
	qpShuffle := q.Get("shuffle")
	if qpShuffle != "" {
		b, err := strconv.ParseBool(qpShuffle)
		if err != nil {
			log.Println(err)
			errorResponse(w, errs.NewBadRequest(fmt.Sprintf("invalid value for shuffle: %s", qpShuffle)))
			return
		}
		shuffle = b
	}

	var cardCodes []string
	qpCards := q.Get("cards")
	if qpCards != "" {
		cardCodes = strings.Split(qpCards, ",")
	}

	deck, err := h.deckService.CreateDeck(shuffle, cardCodes)
	if err != nil {
		log.Println(err)
		errorResponse(w, err)
		return
	}

	jsonResponse(w, deck, http.StatusCreated)
}

// GetDeckHandler handles: GET /decks/{id}.
type GetDeckHandler struct {
	deckService deck.Servicer
}

func NewGetDeckHandler(deckService deck.Servicer) *GetDeckHandler {
	return &GetDeckHandler{deckService}
}

func (h *GetDeckHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	deckId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		errorResponse(w, errs.NewBadRequest(fmt.Sprintf("invalid deck id: %s", id)))
		return
	}

	deck, err := h.deckService.GetDeck(deckId)
	if err != nil {
		log.Println(err)
		errorResponse(w, err)
		return
	}

	jsonResponse(w, deck, http.StatusOK)
}

// DeleteCardsHandler handles: DELETE /decks/{id}/cards?count={int}.
type DeleteCardsHandler struct {
	deckService deck.Servicer
}

func NewDeleteCardsHandler(deckService deck.Servicer) *DeleteCardsHandler {
	return &DeleteCardsHandler{deckService}
}

func (h *DeleteCardsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := r.URL.Query()

	id := vars["id"]
	deckId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		errorResponse(w, errs.NewBadRequest(fmt.Sprintf("invalid deck id: %s", id)))
		return
	}

	count := 1
	qpCount := q.Get("count")
	if qpCount != "" {
		n, err := strconv.Atoi(qpCount)
		if err != nil {
			log.Println(err)
			errorResponse(w, errs.NewBadRequest(fmt.Sprintf("invalid value for count: %s", qpCount)))
			return
		}
		count = n
	}

	deck, err := h.deckService.DrawCards(deckId, count)
	if err != nil {
		log.Println(err)
		errorResponse(w, err)
		return
	}

	jsonResponse(w, deck, http.StatusOK)
}
