package responses

import "github.com/danielgz405/Resev/models"

type TableResponse struct {
	Quantity int            `json:"quantity"`
	Table    []models.Table `json:"table"`
}
