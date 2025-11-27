package handler

import (
	"aggregator/models"
	"aggregator/service"
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func FlightHandler(fs *service.FlightService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mode := models.SortMode(r.URL.Query().Get("sort_by"))
		if !validateSortMode(mode) {
			mode = models.SortByPrice
		}

		flights, err := fs.GetFlights(mode)
		if err != nil {
			http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(flights); err != nil {
			http.Error(w, "Failed to encode flights", http.StatusInternalServerError)
		}
	}
}

func validateSortMode(mode models.SortMode) bool {
	switch mode {
	case models.SortByPrice, models.SortByTime, models.SortByDepartureDate:
		return true
	default:
		return false
	}
}
