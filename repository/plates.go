package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

func InsertPlate(ctx context.Context, typeClient *models.InsertPlate) (*models.Plate, error) {
	return implementation.InsertPlate(ctx, typeClient)
}

func GetPlateById(ctx context.Context, id string) (*models.Plate, error) {
	return implementation.GetPlateById(ctx, id)
}

func ListPlates(ctx context.Context) ([]models.Plate, error) {
	return implementation.ListPlates(ctx)
}

func UpdatePlate(ctx context.Context, data *models.UpdatePlate, id string) (*models.Plate, error) {
	return implementation.UpdatePlate(ctx, data, id)
}

func DeletePlate(ctx context.Context, id string) error {
	return implementation.DeletePlate(ctx, id)
}

func ListPlatesByPage(ctx context.Context, limit int, page int) ([]models.Plate, int, error) {
	return implementation.ListPlatesByPage(ctx, limit, page)
}

func GetPlatesByIds(ctx context.Context, ids []string) ([]models.Plate, error) {
	return implementation.GetPlatesByIds(ctx, ids)
}
