package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

func InsertRole(ctx context.Context, typeClient *models.InsertRole) (*models.Role, error) {
	return implementation.InsertRole(ctx, typeClient)
}

func GetRoleById(ctx context.Context, id string) (*models.Role, error) {
	return implementation.GetRoleById(ctx, id)
}

func ListRoles(ctx context.Context) ([]models.Role, error) {
	return implementation.ListRoles(ctx)
}

func UpdateRole(ctx context.Context, data *models.UpdateRole, id string) (*models.Role, error) {
	return implementation.UpdateRole(ctx, data, id)
}

func DeleteRole(ctx context.Context, id string) error {
	return implementation.DeleteRole(ctx, id)
}

func ListRolesByPage(ctx context.Context, limit int, page int) ([]models.Role, int, error) {
	return implementation.ListRolesByPage(ctx, limit, page)
}
