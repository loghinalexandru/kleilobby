package dst

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/loghinalexandru/klei-lobby/caching"
	"github.com/loghinalexandru/klei-lobby/dst/models"
)

type service struct {
	logger *log.Logger
	cache  *caching.Cache[models.ViewModel]
}

func (s service) GetByServerNameAndHost(token string, region string, serverName string, hostKU string) (models.ViewModel, error) {
	key := fmt.Sprintf("%v_%v_%v", region, serverName, hostKU)
	inCache := s.cache.Contains(key) && time.Now().UTC().Before(s.cache.GetTimestamp(key).Add(1*time.Minute))

	if inCache {
		return s.cache.Get(key), nil
	}

	kleiRequest, err := http.NewRequest("GET", fmt.Sprintf("https://lobby-v2-cdn.klei.com/%v-Steam.json.gz", region), nil)

	if err != nil {
		s.logger.Println(err)
		return models.ViewModel{}, err
	}

	result, err := http.DefaultClient.Do(kleiRequest)

	if err != nil {
		s.logger.Println(err)
		return models.ViewModel{}, err
	}

	content, _ := io.ReadAll(result.Body)
	model := &models.RequestWrapper{}
	json.Unmarshal(content, model)

	for _, server := range model.Lobby {
		if strings.Contains(server.Name, serverName) && server.HostKU == hostKU {
			model, err := s.GetByRowID(server.RowID, token, region)

			if err != nil {
				return models.ViewModel{}, err
			}

			s.cache.Add(key, model)
			return model, nil
		}
	}

	return models.ViewModel{}, errors.New("server not found")
}

func (s service) GetAll(token string, region string) ([]models.ViewModel, error) {
	kleiRequest, err := http.NewRequest("GET", fmt.Sprintf("https://lobby-v2-cdn.klei.com/%v-Steam.json.gz", region), nil)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	result, err := http.DefaultClient.Do(kleiRequest)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	content, _ := io.ReadAll(result.Body)

	model := &models.RequestWrapper{}
	json.Unmarshal(content, model)

	viewModels := make([]models.ViewModel, len(model.Lobby))

	for idx, entry := range model.Lobby {
		mappedEntry, err := MapToViewModel(entry)

		if err != nil {
			s.logger.Println(err)
		}

		viewModels[idx] = mappedEntry
	}

	return viewModels, nil
}

func (s service) GetByRowID(pathRowID string, token string, region string) (models.ViewModel, error) {
	kleiRequest, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://lobby-v2-%v.klei.com/lobby/read", region),
		strings.NewReader(fmt.Sprintf("{\"__gameId\": \"DontStarveTogether\",\"__token\": \"%v\", \"query\":{\"__rowId\":\"%v\"}}}", token, pathRowID)))

	if err != nil {
		s.logger.Println(err)
		return models.ViewModel{}, err
	}

	kleiRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err := http.DefaultClient.Do(kleiRequest)

	if err != nil {
		s.logger.Println(err)
		return models.ViewModel{}, err
	}

	content, _ := io.ReadAll(result.Body)

	model := &models.RequestWrapper{}
	json.Unmarshal(content, model)

	if model == nil || len(model.Lobby) < 1 {
		return models.ViewModel{}, errors.New("missing lobby model")
	}

	data, err := MapToViewModel(model.Lobby[0])

	if err != nil {
		return models.ViewModel{}, err
	}

	return data, nil
}
