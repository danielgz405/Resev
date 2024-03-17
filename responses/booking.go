package responses

import "github.com/danielgz405/Resev/models"

type BookingResponse struct {
	Quantity int              `json:"quantity"`
	Booking  []models.Booking `json:"bookings"`
}
