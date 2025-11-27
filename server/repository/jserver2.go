package repository

import (
	"aggregator/models"
	"encoding/json"
	"os"
	"path/filepath"
)

type JServer2SegmentFlightDTO struct {
	Number string `json:"number"`
	From   string `json:"from"`
	To     string `json:"to"`
	Depart string `json:"depart"`
	Arrive string `json:"arrive"`
}

type JServer2SegmentDTO struct {
	Flight JServer2SegmentFlightDTO `json:"flight"`
}

type JServer2TravelerDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type JServer2TotalDTO struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type JServer2FlightToBookDTO struct {
	Reference string               `json:"reference"`
	Status    string               `json:"status"`
	Traveler  JServer2TravelerDTO  `json:"traveler"`
	Segments  []JServer2SegmentDTO `json:"segments"`
	Total     JServer2TotalDTO     `json:"total"`
}

type JServer2Response struct {
	FlightsToBook []JServer2FlightToBookDTO `json:"flight_to_book"`
}

type JServer2Repository struct {
	FilePath string
}

func NewJServer2Repository(filePath string) *JServer2Repository {
	return &JServer2Repository{FilePath: filePath}
}

func (r *JServer2Repository) GetFlights() ([]models.Flight, error) {
	data, err := os.ReadFile(filepath.Clean(r.FilePath))
	if err != nil {
		return nil, err
	}

	var response JServer2Response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	flights := make([]models.Flight, len(response.FlightsToBook))
	for i, dto := range response.FlightsToBook {
		if len(dto.Segments) == 0 {
			continue
		}

		firstSegment := dto.Segments[0].Flight
		lastSegment := dto.Segments[len(dto.Segments)-1].Flight

		segments := make([]models.Segment, len(dto.Segments))
		for j, segDTO := range dto.Segments {
			segments[j] = models.Segment{
				FlightNumber: segDTO.Flight.Number,
				DepartureTime:    segDTO.Flight.Depart,
				ArrivalTime:      segDTO.Flight.Arrive,
				DepartureAirport:  segDTO.Flight.From,
				ArrivalAirport:    segDTO.Flight.To,
			}
		}

		flights[i] = models.Flight{
			BookingID: dto.Reference,
			Status:    dto.Status,
			Passenger: models.Passenger{
				FirstName: dto.Traveler.FirstName,
				LastName:  dto.Traveler.LastName,
			},
			Segments:         segments,
			FlightNumber:     firstSegment.Number,
			Price:            dto.Total.Amount,
			Currency:         dto.Total.Currency,
			DepartureTime:    firstSegment.Depart,
			ArrivalTime:      lastSegment.Arrive,
			DepartureAirport: firstSegment.From,
			ArrivalAirport:   lastSegment.To,
			TotalTravelTime:  calculateDurationFromString(firstSegment.Depart, lastSegment.Arrive),
			Source:           "JServer2",
		}
	}

	return flights, nil
}