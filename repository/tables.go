package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

func InsertTable(ctx context.Context, typeClient *models.InsertTable) (*models.Table, error) {
	return implementation.InsertTable(ctx, typeClient)
}

func GetTableById(ctx context.Context, id string) (*models.Table, error) {
	return implementation.GetTableById(ctx, id)
}

func ListTables(ctx context.Context) ([]models.Table, error) {
	return implementation.ListTables(ctx)
}

func UpdateTable(ctx context.Context, data *models.UpdateTable, id string) (*models.Table, error) {
	return implementation.UpdateTable(ctx, data, id)
}

func DeleteTable(ctx context.Context, id string) error {
	return implementation.DeleteTable(ctx, id)
}

func ListTablesByPage(ctx context.Context, limit int, page int) ([]models.Table, int, error) {
	return implementation.ListTablesByPage(ctx, limit, page)
}
