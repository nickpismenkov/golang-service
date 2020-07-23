package handler

import (
	"fmt"
	"net/http"
)

// Handler ...
type Handler struct {
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

// Hello ...
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(w, "hello")
}
