package sort

import (
	"aggregator/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestByPrice(t *testing.T) {
	flights := []models.Flight{
		{Price: 1000.0},
		{Price: 500.0},
		{Price: 750.0},
	}

	ByPrice(flights)

	assert.Equal(t, 500.0, flights[0].Price)
	assert.Equal(t, 750.0, flights[1].Price)
	assert.Equal(t, 1000.0, flights[2].Price)
}

func TestByTravelTime(t *testing.T) {
	flights := []models.Flight{
		{TotalTravelTime: "2h30m"},
		{TotalTravelTime: "1h15m"},
		{TotalTravelTime: "3h45m"},
	}

	ByTravelTime(flights)

	assert.Equal(t, "1h15m", flights[0].TotalTravelTime)
	assert.Equal(t, "2h30m", flights[1].TotalTravelTime)
	assert.Equal(t, "3h45m", flights[2].TotalTravelTime)
}

func TestByDepartureDate(t *testing.T) {
	flights := []models.Flight{
		{DepartureTime: "2026-01-03T10:00:00Z"},
		{DepartureTime: "2026-01-01T10:00:00Z"},
		{DepartureTime: "2026-01-02T10:00:00Z"},
	}

	ByDepartureDate(flights)

	time1, _ := time.Parse(time.RFC3339, flights[0].DepartureTime)
	time2, _ := time.Parse(time.RFC3339, flights[1].DepartureTime)
	time3, _ := time.Parse(time.RFC3339, flights[2].DepartureTime)

	assert.True(t, time1.Before(time2))
	assert.True(t, time2.Before(time3))
}

func TestByPrice_EmptySlice(t *testing.T) {
	flights := []models.Flight{}
	ByPrice(flights)
	assert.Empty(t, flights)
}

func TestByTravelTime_InvalidDuration(t *testing.T) {
	flights := []models.Flight{
		{TotalTravelTime: "invalid"},
		{TotalTravelTime: "1h"},
	}

	ByTravelTime(flights)

	assert.Equal(t, "invalid", flights[0].TotalTravelTime)
	assert.Equal(t, "1h", flights[1].TotalTravelTime)
}

func TestByDepartureDate_InvalidTime(t *testing.T) {
	flights := []models.Flight{
		{DepartureTime: "invalid"},
		{DepartureTime: "2026-01-01T10:00:00Z"},
	}

	ByDepartureDate(flights)

	assert.Equal(t, "2026-01-01T10:00:00Z", flights[0].DepartureTime)
	assert.Equal(t, "invalid", flights[1].DepartureTime)
}

