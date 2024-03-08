package repository

import (
	"context"

	"github.com/danielgz405/Resev/models"
)

type Repository interface {
	//Users
	InsertUser(ctx context.Context, user *models.InsertUser) (*models.Profile, error)
	GetUserById(ctx context.Context, id string) (*models.Profile, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, data models.UpdateUser) (*models.Profile, error)
	DeleteUser(ctx context.Context, id string) error

	//tables
	InsertTable(ctx context.Context, typeClient *models.InsertTable) (*models.Table, error)
	GetTableById(ctx context.Context, id string) (*models.Table, error)
	ListTables(ctx context.Context) ([]models.Table, error)
	UpdateTable(ctx context.Context, data *models.UpdateTable, id string) (*models.Table, error)
	DeleteTable(ctx context.Context, id string) error
	ListTablesByPage(ctx context.Context, limit int, page int) ([]models.Table, int, error)

	//roles
	InsertRole(ctx context.Context, typeClient *models.InsertRole) (*models.Role, error)
	GetRoleById(ctx context.Context, id string) (*models.Role, error)
	ListRoles(ctx context.Context) ([]models.Role, error)
	UpdateRole(ctx context.Context, data *models.UpdateRole, id string) (*models.Role, error)
	DeleteRole(ctx context.Context, id string) error
	ListRolesByPage(ctx context.Context, limit int, page int) ([]models.Role, int, error)

	//plates
	InsertPlate(ctx context.Context, typeClient *models.InsertPlate) (*models.Plate, error)
	GetPlateById(ctx context.Context, id string) (*models.Plate, error)
	ListPlates(ctx context.Context) ([]models.Plate, error)
	UpdatePlate(ctx context.Context, data *models.UpdatePlate, id string) (*models.Plate, error)
	DeletePlate(ctx context.Context, id string) error
	ListPlatesByPage(ctx context.Context, limit int, page int) ([]models.Plate, int, error)
	GetPlatesByIds(ctx context.Context, ids []string) ([]models.Plate, error)

	//Close the connection
	Close() error
}

var implementation Repository

// Repo
func SetRepository(repository Repository) {
	implementation = repository
}

// Close the connection
func Close() error {
	return implementation.Close()
}
