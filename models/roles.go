package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Id
//Name
//Description

type Role struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

type InsertRole struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

type UpdateRole struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}
