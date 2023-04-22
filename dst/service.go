package dst

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/loghinalexandru/kleilobby/caching"
	"github.com/loghinalexandru/kleilobby/dst/model"
)

var (
	ErrNotFound = errors.New("resource not found")
)

const (
	lobbyURL     = "https://lobby-v2-cdn.klei.com/%v-Steam.json.gz"
	lobbyReadURL = "https://lobby-v2-%v.klei.com/lobby/read"
)

type service struct {
	client *http.Client
	logger *log.Logger
	cache  *caching.Cache[model.ViewModel]
}

func (s service) GetByServerNameAndHost(token string, region string, serverName string, hostKU string) (model.ViewModel, error) {
	key := fmt.Sprintf("%v_%v_%v", region, serverName, hostKU)

	if s.cache.Contains(key) {
		s.logger.Printf("cache hit for key: %v", key)
		return s.cache.Get(key), nil
	}

	request, err := http.NewRequest("GET", fmt.Sprintf(lobbyURL, region), nil)

	if err != nil {
		s.logger.Println(err)
		return model.ViewModel{}, err
	}

	result, err := s.client.Do(request)

	if err != nil {
		s.logger.Println(err)
		return model.ViewModel{}, err
	}

	content, _ := io.ReadAll(result.Body)
	wrapper := &model.RequestWrapper{}
	json.Unmarshal(content, wrapper)

	for _, server := range wrapper.Lobby {
		if strings.Contains(server.Name, serverName) && server.HostKU == hostKU {
			viewModel, err := s.GetByRowID(token, region, server.RowID)

			if err != nil {
				return model.ViewModel{}, err
			}

			s.cache.Add(key, viewModel)
			return viewModel, nil
		}
	}

	return model.ViewModel{}, ErrNotFound
}

func (s service) GetAll(region string) ([]model.ViewModel, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(lobbyURL, region), nil)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	result, err := s.client.Do(request)

	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	content, _ := io.ReadAll(result.Body)

	wrapper := &model.RequestWrapper{}
	json.Unmarshal(content, wrapper)

	viewModels := make([]model.ViewModel, len(wrapper.Lobby))

	for i, entry := range wrapper.Lobby {
		mappedEntry, err := MapToViewModel(entry)

		if err != nil {
			s.logger.Println(err)
		}

		viewModels[i] = mappedEntry
	}

	return viewModels, nil
}

func (s service) GetByRowID(token string, region string, pathRowID string) (model.ViewModel, error) {
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(lobbyReadURL, region),
		strings.NewReader(fmt.Sprintf("{\"__gameId\": \"DontStarveTogether\",\"__token\": \"%v\", \"query\":{\"__rowId\":\"%v\"}}}", token, pathRowID)))

	if err != nil {
		s.logger.Println(err)
		return model.ViewModel{}, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err := s.client.Do(request)

	if err != nil {
		s.logger.Println(err)
		return model.ViewModel{}, err
	}

	content, _ := io.ReadAll(result.Body)

	wrapper := &model.RequestWrapper{}
	json.Unmarshal(content, wrapper)

	if wrapper == nil || len(wrapper.Lobby) < 1 {
		return model.ViewModel{}, ErrNotFound
	}

	data, err := MapToViewModel(wrapper.Lobby[0])

	if err != nil {
		return model.ViewModel{}, err
	}

	return data, nil
}
