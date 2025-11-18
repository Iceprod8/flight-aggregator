package service

import (
	"aggregator/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFlightRepository struct {
	mock.Mock
}

func (m *MockFlightRepository) GetFlights() ([]models.Flight, error) {
	args := m.Called()
	return args.Get(0).([]models.Flight), args.Error(1)
}

func TestFlightService_GetFlights_ByPrice(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	flights1 := []models.Flight{
		{Price: 1000.0, BookingID: "A1"},
		{Price: 500.0, BookingID: "A2"},
	}

	flights2 := []models.Flight{
		{Price: 750.0, BookingID: "B1"},
		{Price: 300.0, BookingID: "B2"},
	}

	mockRepo1.On("GetFlights").Return(flights1, nil)
	mockRepo2.On("GetFlights").Return(flights2, nil)

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByPrice)

	assert.NoError(t, err)
	assert.Len(t, result, 4)
	assert.Equal(t, 300.0, result[0].Price)
	assert.Equal(t, 500.0, result[1].Price)
	assert.Equal(t, 750.0, result[2].Price)
	assert.Equal(t, 1000.0, result[3].Price)

	mockRepo1.AssertExpectations(t)
	mockRepo2.AssertExpectations(t)
}

func TestFlightService_GetFlights_ByTime(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	flights1 := []models.Flight{
		{TotalTravelTime: "3h", BookingID: "A1"},
		{TotalTravelTime: "1h", BookingID: "A2"},
	}

	flights2 := []models.Flight{
		{TotalTravelTime: "2h", BookingID: "B1"},
	}

	mockRepo1.On("GetFlights").Return(flights1, nil)
	mockRepo2.On("GetFlights").Return(flights2, nil)

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByTime)

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "1h", result[0].TotalTravelTime)
	assert.Equal(t, "2h", result[1].TotalTravelTime)
	assert.Equal(t, "3h", result[2].TotalTravelTime)
}

func TestFlightService_GetFlights_ByDepartureDate(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	flights1 := []models.Flight{
		{DepartureTime: "2026-01-03T10:00:00Z", BookingID: "A1"},
		{DepartureTime: "2026-01-01T10:00:00Z", BookingID: "A2"},
	}

	flights2 := []models.Flight{
		{DepartureTime: "2026-01-02T10:00:00Z", BookingID: "B1"},
	}

	mockRepo1.On("GetFlights").Return(flights1, nil)
	mockRepo2.On("GetFlights").Return(flights2, nil)

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByDepartureDate)

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "2026-01-01T10:00:00Z", result[0].DepartureTime)
	assert.Equal(t, "2026-01-02T10:00:00Z", result[1].DepartureTime)
	assert.Equal(t, "2026-01-03T10:00:00Z", result[2].DepartureTime)
}

func TestFlightService_GetFlights_DefaultSort(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	flights1 := []models.Flight{
		{Price: 1000.0, BookingID: "A1"},
	}

	flights2 := []models.Flight{
		{Price: 500.0, BookingID: "B1"},
	}

	mockRepo1.On("GetFlights").Return(flights1, nil)
	mockRepo2.On("GetFlights").Return(flights2, nil)

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortMode("unknown"))

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 500.0, result[0].Price)
	assert.Equal(t, 1000.0, result[1].Price)
}

func TestFlightService_GetFlights_Repo1Error(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	mockRepo1.On("GetFlights").Return([]models.Flight{}, errors.New("repo1 error"))

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByPrice)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "repo1 error", err.Error())
}

func TestFlightService_GetFlights_Repo2Error(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	flights1 := []models.Flight{
		{Price: 1000.0, BookingID: "A1"},
	}

	mockRepo1.On("GetFlights").Return(flights1, nil)
	mockRepo2.On("GetFlights").Return([]models.Flight{}, errors.New("repo2 error"))

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByPrice)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "repo2 error", err.Error())
}

func TestFlightService_GetFlights_EmptyResults(t *testing.T) {
	mockRepo1 := new(MockFlightRepository)
	mockRepo2 := new(MockFlightRepository)

	mockRepo1.On("GetFlights").Return([]models.Flight{}, nil)
	mockRepo2.On("GetFlights").Return([]models.Flight{}, nil)

	service := NewFlightService(mockRepo1, mockRepo2)
	result, err := service.GetFlights(models.SortByPrice)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

