package models

type Segment struct {
	FlightNumber     string `json:"flight_number"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalAirport   string `json:"arrival_airport"`
	DepartureTime    string `json:"departure_time"`
	ArrivalTime      string `json:"arrival_time"`
	Duration         string `json:"duration"`
}
