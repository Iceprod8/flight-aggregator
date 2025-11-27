package repository

import (
	"strings"
	"time"
)

func calculateDuration(departTime, arriveTime time.Time) string {
	if departTime.IsZero() || arriveTime.IsZero() {
		return "N/A"
	}
	return arriveTime.Sub(departTime).String()
}

func calculateDurationFromString(departStr, arriveStr string) string {
    depart, errD := time.Parse(time.RFC3339, departStr)
    arrive, errA := time.Parse(time.RFC3339, arriveStr)
    if errD != nil || errA != nil {
        return "N/A"
    }
    return arrive.Sub(depart).String()
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Fields(fullName)
	if len(parts) > 0 {
		firstName = parts[0]
	}
	if len(parts) > 1 {
		lastName = strings.Join(parts[1:], " ")
	}
	return firstName, lastName
}