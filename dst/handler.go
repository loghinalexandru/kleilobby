package dst

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/loghinalexandru/klei-lobby/caching"
	"github.com/loghinalexandru/klei-lobby/dst/models"
)

type Handler struct {
	logger *log.Logger
	svc    service
}

func NewHandler(l *log.Logger) *Handler {
	cache := caching.New[models.ViewModel]()

	return &Handler{
		logger: l,
		svc: service{
			logger: l,
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
