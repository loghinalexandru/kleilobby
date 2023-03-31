package dst

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/loghinalexandru/klei-lobby/caching"
	"github.com/loghinalexandru/klei-lobby/dst/models"
)

type Handler struct {
	logger *log.Logger
	svc    service
}

func NewHandler(log *log.Logger) *Handler {
	cache := caching.New[models.ViewModel](1 * time.Minute)

	return &Handler{
		logger: log,
		svc: service{
			logger: log,
			cache:  cache,
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

func (h *Handler) ServerName(writer http.ResponseWriter, request *http.Request, serverName string, hostKU string) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result, err := h.svc.GetByServerNameAndHost(request.URL.Query().Get("token"), request.URL.Query().Get("region"), serverName, hostKU)

	if errors.Is(err, ErrNotFound) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

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

func (h *Handler) RowID(writer http.ResponseWriter, request *http.Request, pathRowID string) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result, err := h.svc.GetByRowID(pathRowID, request.URL.Query().Get("token"), request.URL.Query().Get("region"))

	if errors.Is(err, ErrNotFound) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

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
