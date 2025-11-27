package repository

import (
	"aggregator/models"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type FlightsDTO struct {
	BookingId string `json:"bookingId"`
	Status string `json:"status"`
    PassengerName string `json:"passengerName"`
    FlightNumber string `json:"flightNumber"`
    DepartureAirport string `json:"departureAirport"`
    ArrivalAirport string `json:"arrivalAirport"`
    DepartureTime time.Time `json:"departureTime"`
    ArrivalTime time.Time `json:"arrivalTime"`
	Price float64 `json:"price"`
    Currency string `json:"currency"`
    Id string `json:"id"`
}

type JServer1Response struct {
	Flights []FlightsDTO `json:"flights"`
}

type JServer1Repository struct {
	FilePath string
}

func NewJServer1Repository(filePath string) *JServer1Repository {
	return &JServer1Repository{FilePath: filePath}
}

func (r *JServer1Repository) GetFlights() ([]models.Flight, error) {
	data, err := os.ReadFile(filepath.Clean(r.FilePath))
	if err != nil {
		return nil, err
	}

	var response JServer1Response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	flights := make([]models.Flight, len(response.Flights))
	for i, dto := range response.Flights {
		firstName, lastName := splitName(dto.PassengerName)

		flights[i] = models.Flight{
			BookingID:      dto.BookingId,
			Status:         dto.Status,
			Passenger: models.Passenger{
				FirstName: firstName,
				LastName:  lastName,
			},
			FlightNumber:   dto.FlightNumber,
			Price:          float64(dto.Price),
			Currency:       dto.Currency,
			DepartureTime:  dto.DepartureTime.Format("2006-01-02T15:04:05Z"),
			ArrivalTime:    dto.ArrivalTime.Format("2006-01-02T15:04:05Z"),
			DepartureAirport: dto.DepartureAirport,
			ArrivalAirport: dto.ArrivalAirport,
			TotalTravelTime: calculateDuration(dto.DepartureTime, dto.ArrivalTime),
			Source:         "JServer1",
		}
	}
	return flights, nil
}
