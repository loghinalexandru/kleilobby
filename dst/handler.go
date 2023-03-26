package dst

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	logger *log.Logger
	svc    service
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{
		logger: l,
		svc: service{
			logger: l,
		},
	}
}

func (h *Handler) All(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result, err := h.svc.GetAll(request.URL.Query().Get("token"), request.URL.Query().Get("region"))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(data)
}

func (h *Handler) ServerName(writer http.ResponseWriter, request *http.Request) {
}

func (h *Handler) RowID(writer http.ResponseWriter, request *http.Request, pathRowID string) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result, err := h.svc.GetByRowID(pathRowID, request.URL.Query().Get("token"), request.URL.Query().Get("region"))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(data)
}
