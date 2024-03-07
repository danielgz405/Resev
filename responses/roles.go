package responses

import "github.com/danielgz405/Resev/models"

type RoleResponse struct {
	Quantity int           `json:"quantity"`
	Role     []models.Role `json:"role"`
}
