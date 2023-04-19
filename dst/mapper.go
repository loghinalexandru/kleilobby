package dst

import (
	"regexp"
	"strconv"

	"github.com/loghinalexandru/klei-lobby/dst/model"
)

func MapToViewModel(input model.ServerInfo) (model.ViewModel, error) {
	dayMatch := regexp.MustCompile("day=([0-9]+)")
	daysElapsedInSeasonMatch := regexp.MustCompile("dayselapsedinseason=([0-9]+)")
	daysLeftInSeasonMatch := regexp.MustCompile("daysleftinseason=([0-9]+)")

	var day, daysElapsedInSeason, daysLeftInSeason int

	if dayMatch.MatchString(input.RawServerData) {
		day, _ = strconv.Atoi(dayMatch.FindStringSubmatch(input.RawServerData)[1])
	}

	if daysElapsedInSeasonMatch.MatchString(input.RawServerData) {
		daysElapsedInSeason, _ = strconv.Atoi(daysElapsedInSeasonMatch.FindStringSubmatch(input.RawServerData)[1])
	}

	if daysLeftInSeasonMatch.MatchString(input.RawServerData) {
		daysLeftInSeason, _ = strconv.Atoi(daysLeftInSeasonMatch.FindStringSubmatch(input.RawServerData)[1])
	}

	data := model.ViewModel{
		Address:          input.Address,
		RowID:            input.RowID,
		HostKU:           input.HostKU,
		ServerName:       input.Name,
		UsesModes:        input.Mods,
		PlayerCount:      input.Connected,
		Season:           input.Season,
		Day:              day,
		DaysInSeason:     daysElapsedInSeason,
		DaysLeftInSeason: daysLeftInSeason,
	}

	return data, nil
}
