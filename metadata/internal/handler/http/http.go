package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/stevan1008/movieGolang/metadata/internal/controller/metadata"
	"github.com/stevan1008/movieGolang/metadata/internal/repository"
)

type Handler struct {
	ctlr *metadata.Controller
}

func New(ctlr *metadata.Controller) *Handler {
	return &Handler{ctlr: ctlr}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	m, err := h.ctlr.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error for movie %s: %v/n", id, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v/n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
