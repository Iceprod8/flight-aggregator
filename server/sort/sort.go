package sort

import (
	"aggregator/models"
	"sort"
	"time"
)

func ByPrice(flights []models.Flight) {
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})
}

func ByTravelTime(flights []models.Flight) {
	sort.Slice(flights, func(i, j int) bool {
		durationI := parseDuration(flights[i].TotalTravelTime)
		durationJ := parseDuration(flights[j].TotalTravelTime)
		return durationI < durationJ
	})
}

func ByDepartureDate(flights []models.Flight) {
	sort.Slice(flights, func(i, j int) bool {
		timeI, errI := time.Parse(time.RFC3339, flights[i].DepartureTime)
		timeJ, errJ := time.Parse(time.RFC3339, flights[j].DepartureTime)
		
		if errI != nil {
			return false
		}
		if errJ != nil {
			return true
		}
		
		return timeI.Before(timeJ)
	})
}

func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0
	}
	return duration
}

