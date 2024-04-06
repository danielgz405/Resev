package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Phone    string             `bson:"phone" json:"phone"`
	Role_id  string             `bson:"role_id" json:"role_id"`
	Image    string             `bson:"image" json:"image"`
	ImageRef string             `bson:"image_ref" json:"image_ref"`
	Bookings []string           `bson:"bookings" json:"bookings"`
}

type Profile struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Phone    string             `bson:"phone" json:"phone"`
	Role_id  string             `bson:"role_id" json:"role_id"`
	Image    string             `bson:"image" json:"image"`
	ImageRef string             `bson:"image_ref" json:"image_ref"`
	Bookings []string           `bson:"bookings" json:"bookings"`
}

type InsertUser struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Phone    string `bson:"phone" json:"phone"`
	Role_id  string `bson:"role_id" json:"role_id"`
	Password string `bson:"password" json:"password"`
}

type UpdateUser struct {
	Id       string   `bson:"_id" json:"_id"`
	Name     string   `bson:"name" json:"name"`
	Email    string   `bson:"email" json:"email"`
	Phone    string   `bson:"phone" json:"phone"`
	Role_id  string   `bson:"role_id" json:"role_id"`
	Image    string   `bson:"image" json:"image"`
	ImageRef string   `bson:"image_ref" json:"image_ref"`
	Bookings []string `bson:"bookings" json:"bookings"`
}
