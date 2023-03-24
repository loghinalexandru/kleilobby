package dst

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/loghinalexandru/klei-lobby/models"
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
	mux.HandleFunc("/rowid", h.RowId)
}

func (h *Handler) All(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) RowId(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)

	payload := fmt.Sprintf("{\"__gameId\": \"DontStarveTogether\",\"__token\": \"%v\", \"query\":{\"__rowId\":\"%v\"}}}", request.URL.Query().Get("token"), request.URL.Path)

	kleiRequest, err := http.NewRequest(
		"POST",
		"https://lobby-v2-eu-central-1.klei.com/lobby/read",
		strings.NewReader(payload))

	if err != nil {
		h.logger.Println(err)
	}

	kleiRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result, err := http.DefaultClient.Do(kleiRequest)

	if err != nil {
		h.logger.Println(err)
	}

	content, _ := io.ReadAll(result.Body)
	model := &models.RequestWrapper{}

	json.Unmarshal(content, model)
}
