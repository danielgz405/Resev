package responses

import "github.com/danielgz405/Resev/models"

type UserResponse struct {
	Id       string           `bson:"_id" json:"_id"`
	Name     string           `bson:"name" json:"name"`
	Email    string           `bson:"email" json:"email"`
	Phone    string           `bson:"phone" json:"phone"`
	Role     models.Role      `bson:"role" json:"role"`
	Image    string           `bson:"image" json:"image"`
	ImageRef string           `bson:"image_ref" json:"image_ref"`
	Bookings []models.Booking `json:"bookings"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
