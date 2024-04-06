package responses

import (
	"time"

	"github.com/danielgz405/Resev/models"
)

type BookingsResponse struct {
	Quantity int               `json:"quantity"`
	Booking  []BookingResponse `json:"bookings"`
}

type BookingResponse struct {
	Id          string         `json:"_id"`
	Table       models.Table   `json:"table"`
	User        models.Profile `json:"user"`
	Order       models.Order   `json:"order"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	Hour        string         `json:"hour"`
}
