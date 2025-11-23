package handler

import (
	"aggregator/models"
	"aggregator/service"
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func FlightHandler(fs *service.FlightService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sortParams := r.URL.Query().Get("sort_by")
		var mode models.SortMode
		if !validateSortMode(mode) {
			mode = models.SortByPrice
		} else {
			mode = models.SortMode(sortParams)
		}

		flights, err := fs.GetFlights(mode)
		if err != nil {
			http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Résultats = %v", flights)))
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
