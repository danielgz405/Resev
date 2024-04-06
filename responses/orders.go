package responses

import "github.com/danielgz405/Resev/models"

type OrdersResponse struct {
	Quantity int             `json:"quantity"`
	Order    []OrderResponse `json:"order"`
}

type OrderResponse struct {
	Id        string         `json:"_id"`
	Table     models.Table   `bjson:"table"`
	User      models.Profile `json:"user"`
	SubTotal  float64        `json:"subtotal"`
	Iva       float64        `json:"iva"`
	Total     float64        `json:"total"`
	TimeStamp float64        `json:"timeStamp"`
	Plates    []models.Plate `json:"plates"`
}
