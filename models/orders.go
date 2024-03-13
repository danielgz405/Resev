package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id        primitive.ObjectID   `bson:"_id" json:"_id"`
	TableId   string               `bson:"table_id" json:"table_id"`
	UserId    string               `bson:"user_id" json:"user_id"`
	SubTotal  float64              `bson:"subtotal" json:"subtotal"`
	Iva       float64              `bson:"iva" json:"iva"`
	Total     float64              `bson:"total" json:"total"`
	TimeStamp float64              `bson:"timeStamp" json:"timeStamp"`
	PlatesId  []primitive.ObjectID `bson:"plates_id" json:"plates_id"`
}

type InsertOrder struct {
	TableId  string               `bson:"table_id" json:"table_id"`
	UserId   string               `bson:"user_id" json:"user_id"`
	SubTotal float64              `bson:"subtotal" json:"subtotal"`
	Iva      float64              `bson:"iva" json:"iva"`
	Total    float64              `bson:"total" json:"total"`
	PlatesId []primitive.ObjectID `bson:"plates_id" json:"plates_id"`
}

type UpdateOrder struct {
	TableId   string               `bson:"table_id" json:"table_id"`
	SubTotal  float64              `bson:"subtotal" json:"subtotal"`
	Total     float64              `bson:"total" json:"total"`
	TimeStamp float64              `bson:"timeStamp" json:"timeStamp"`
	PlatesId  []primitive.ObjectID `bson:"plates_id" json:"plates_id"`
}
