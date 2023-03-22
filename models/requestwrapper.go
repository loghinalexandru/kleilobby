package models

type RequestWrapper struct {
	Lobby []ServerInfo `json:"GET"`
}
