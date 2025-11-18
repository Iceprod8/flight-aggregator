package repository

import "aggregator/models"

type FlightRepository interface {
	GetFlights() ([]models.Flight, error)
}

