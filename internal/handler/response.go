package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gkatanacio/card-deck-api/internal/errs"
)

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(data)
}

type errorResponseBody struct {
	Error string `json:"error"`
}

func errorResponse(w http.ResponseWriter, err error) {
	body := &errorResponseBody{}
	body.Error = err.Error()

	var status int
	var httpErr errs.HttpError
	switch {
	case errors.As(err, &httpErr):
		status = httpErr.StatusCode()
	default:
		status = http.StatusInternalServerError
	}

	jsonResponse(w, body, status)
}
