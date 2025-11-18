package service

import (
	"aggregator/models"
	"aggregator/repository"
	"aggregator/sort"
)

type FlightService struct {
	repo1 repository.FlightRepository
	repo2 repository.FlightRepository
}

func NewFlightService(repo1, repo2 repository.FlightRepository) *FlightService {
	return &FlightService{
		repo1: repo1,
		repo2: repo2,
	}
}

func (s *FlightService) GetFlights(sortMode models.SortMode) ([]models.Flight, error) {
	flights1, err := s.repo1.GetFlights()
	if err != nil {
		return nil, err
	}

	flights2, err := s.repo2.GetFlights()
	if err != nil {
		return nil, err
	}

	allFlights := append(flights1, flights2...)

	switch sortMode {
	case models.SortByPrice:
		sort.ByPrice(allFlights)
	case models.SortByTime:
		sort.ByTravelTime(allFlights)
	case models.SortByDepartureDate:
		sort.ByDepartureDate(allFlights)
	default:
		sort.ByPrice(allFlights)
	}

	return allFlights, nil
}

