package responses

import "github.com/danielgz405/Resev/models"

type OrderResponse struct {
	Quantity int            `json:"quantity"`
	Order    []models.Order `json:"order"`
}
