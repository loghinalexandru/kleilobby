package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/loghinalexandru/klei-lobby/models"
)

type DontStarveTogether struct {
	logger *log.Logger
}

func NewDontStarveTogether(l *log.Logger) *DontStarveTogether {
	return &DontStarveTogether{
		logger: l,
	}
}

func (h *DontStarveTogether) All(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	writer.WriteHeader(http.StatusOK)
}

func (h *DontStarveTogether) RowId(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Println(request.URL.Path)

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
