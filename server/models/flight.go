package models

type Flight struct {
	// Booking/Reference Information
	BookingID string `json:"booking_id"`
	Status    string `json:"status"`

	// Passenger Information
	Passenger Passenger `json:"passenger"`

	// Flight Details
	FlightNumber string    `json:"flight_number"`
	Segments     []Segment `json:"segments"`

	// Pricing
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`

	// Calculated Fields
	TotalTravelTime  string `json:"total_travel_time"`
	DepartureTime    string `json:"departure_time"`
	ArrivalTime      string `json:"arrival_time"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalAirport   string `json:"arrival_airport"`

	// Metadata
	Source string `json:"source"`
}
