package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	TableId     primitive.ObjectID `bson:"table_id" json:"table_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	OrderId     primitive.ObjectID `bson:"order_id" json:"order_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Date        time.Time          `bson:"date" json:"date"`
	Hour        string             `bson:"hour" json:"hour"`
}

type InsertBooking struct {
	TableId     primitive.ObjectID `bson:"table_id" json:"table_id"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
	OrderId     primitive.ObjectID `bson:"order_id" json:"order_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Date        time.Time          `bson:"date" json:"date"`
	Hour        string             `bson:"hour" json:"hour"`
}

type UpdateBooking struct {
	TableId     primitive.ObjectID `bson:"table_id" json:"table_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Date        time.Time          `bson:"date" json:"date"`
	Hour        string             `bson:"hour" json:"hour"`
}
