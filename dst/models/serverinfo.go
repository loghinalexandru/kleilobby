package models

type ServerInfo struct {
	Address       string `json:"__addr"`
	RowID         string `json:"__rowId"`
	HostKU        string `json:"host"`
	Name          string `json:"name"`
	Password      bool   `json:"password"`
	Mods          bool   `json:"mods"`
	Connected     int    `json:"connected"`
	Season        string `json:"season"`
	ServerPaused  string `json:"serverpaused"`
	RawServerData string `json:"data"`
}
