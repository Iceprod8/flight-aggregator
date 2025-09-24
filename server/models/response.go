package models

type Response struct {
	Flights []Flight `json:"flights"`
	Total   int      `json:"total"`
	SortBy  string   `json:"sort_by,omitempty"`
}
