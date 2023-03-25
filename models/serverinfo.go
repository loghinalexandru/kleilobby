package models

import (
	"regexp"
	"strconv"
)

var (
	dayMatch           = regexp.MustCompile("day=([0-9]+)")
	daysLapsedInSeason = regexp.MustCompile("dayselapsedinseason=([0-9]+)")
	daysLeftInSeason   = regexp.MustCompile("daysleftinseason=([0-9]+)")
)

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

type serverData struct {
	day                int
	daysLapsedInSeason int
	daysLeftInSeason   int
}

// TODO: maybe skip this and use object?
func (s *ServerInfo) NewServerData() *serverData {
	day, _ := strconv.Atoi(dayMatch.FindStringSubmatch(s.RawServerData)[1])
	daysLapsedInSeason, _ := strconv.Atoi(daysLapsedInSeason.FindStringSubmatch(s.RawServerData)[1])
	daysLeftInSeason, _ := strconv.Atoi(daysLeftInSeason.FindStringSubmatch(s.RawServerData)[1])

	return &serverData{
		day:                day,
		daysLapsedInSeason: daysLapsedInSeason,
		daysLeftInSeason:   daysLeftInSeason,
	}
}
