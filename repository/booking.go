package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

func InsertBooking(ctx context.Context, data *models.InsertBooking) (*models.Booking, error) {
	return implementation.InsertBooking(ctx, data)
}

func GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	return implementation.GetBookingById(ctx, id)
}

func ListBookings(ctx context.Context) ([]models.Booking, error) {
	return implementation.ListBookings(ctx)
}

func UpdateBooking(ctx context.Context, data *models.UpdateBooking, id string) (*models.Booking, error) {
	return implementation.UpdateBooking(ctx, data, id)
}

func DeleteBooking(ctx context.Context, id string) error {
	return implementation.DeleteBooking(ctx, id)
}

func ListBookingsByPage(ctx context.Context, limit int, page int) ([]models.Booking, int, error) {
	return implementation.ListBookingsByPage(ctx, limit, page)
}

func GetBookingsByIds(ctx context.Context, ids []string) ([]models.Booking, error) {
	return implementation.GetBookingsByIds(ctx, ids)
}
