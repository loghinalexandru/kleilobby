package dst

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/loghinalexandru/kleilobby/caching"
	"github.com/loghinalexandru/kleilobby/dst/model"
)

const cacheTTL = 5 * time.Minute

type Handler struct {
	logger *slog.Logger
	svc    service
}

func NewHandler(log *slog.Logger) *Handler {
	cache := caching.New[model.ViewModel](cacheTTL)

	return &Handler{
		logger: log,
		svc: service{
			logger: log,
			client: http.DefaultClient,
			cache:  cache,
		},
	}
}

func (h *Handler) All(writer http.ResponseWriter, request *http.Request) {
	result, err := h.svc.GetAll(request.URL.Query().Get("region"))

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
	_, _ = writer.Write(data)
}

func (h *Handler) ServerName(writer http.ResponseWriter, request *http.Request) {
	hostKU := request.PathValue("hostKU")
	serverName := request.PathValue("serverName")

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
	_, _ = writer.Write(data)
}

func (h *Handler) RowID(writer http.ResponseWriter, request *http.Request) {
	rowID := request.PathValue("rowID")
	result, err := h.svc.GetByRowID(request.URL.Query().Get("token"), request.URL.Query().Get("region"), rowID)

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
	_, _ = writer.Write(data)
}
