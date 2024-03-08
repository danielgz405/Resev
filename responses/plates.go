package responses

import "github.com/danielgz405/Resev/models"

type PlateResponse struct {
	Quantity int            `json:"quantity"`
	Plate    []models.Plate `json:"plate"`
}
