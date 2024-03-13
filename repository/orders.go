package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

func InsertOrder(ctx context.Context, typeClient *models.InsertOrder) (*models.Order, error) {
	return implementation.InsertOrder(ctx, typeClient)
}

func GetOrderById(ctx context.Context, id string) (*models.Order, error) {
	return implementation.GetOrderById(ctx, id)
}

func ListOrders(ctx context.Context) ([]models.Order, error) {
	return implementation.ListOrders(ctx)
}

func UpdateOrder(ctx context.Context, data *models.UpdateOrder, id string) (*models.Order, error) {
	return implementation.UpdateOrder(ctx, data, id)
}

func DeleteOrder(ctx context.Context, id string) error {
	return implementation.DeleteOrder(ctx, id)
}

func ListOrdersByPage(ctx context.Context, limit int, page int) ([]models.Order, int, error) {
	return implementation.ListOrdersByPage(ctx, limit, page)
}
