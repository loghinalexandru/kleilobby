package models

const ()

type serverData struct {
	day                 int
	dayseLapsedInSeason int
	daysLeftInSeason    int
}

func New(rawServerData string) *serverData {
	//TODO: build based on raw data
	return &serverData{}
}
