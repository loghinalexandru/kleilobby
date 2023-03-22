package dst

import (
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{
		logger: l,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/all", h.All)
}

func (h *Handler) All(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
