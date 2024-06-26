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
	GetRoleByName(ctx context.Context, name string) (*models.Role, error)
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

	//Otrder
	InsertOrder(ctx context.Context, typeClient *models.InsertOrder) (*models.Order, error)
	GetOrderById(ctx context.Context, id string) (*models.Order, error)
	ListOrders(ctx context.Context) ([]models.Order, error)
	UpdateOrder(ctx context.Context, data *models.UpdateOrder, id string) (*models.Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrdersByPage(ctx context.Context, limit int, page int) ([]models.Order, int, error)

	//boooking
	InsertBooking(ctx context.Context, typeClient *models.InsertBooking) (*models.Booking, error)
	GetBookingById(ctx context.Context, id string) (*models.Booking, error)
	ListBookings(ctx context.Context) ([]models.Booking, error)
	UpdateBooking(ctx context.Context, data *models.UpdateBooking, id string) (*models.Booking, error)
	DeleteBooking(ctx context.Context, id string) error
	ListBookingsByPage(ctx context.Context, limit int, page int) ([]models.Booking, int, error)
	GetBookingsByIds(ctx context.Context, ids []string) ([]models.Booking, error)

	//audit
	AuditOperation(ctx context.Context, id string, table string, operationType string) error

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
