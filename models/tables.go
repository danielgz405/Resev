package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Table struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Number      string             `bson:"number" json:"number"`
	Observation string             `bson:"observation" json:"observation"`
}

type InsertTable struct {
	Number      string `bson:"number" json:"number"`
	Observation string `bson:"observation" json:"observation"`
}

type UpdateTable struct {
	Number      string `bson:"number" json:"number"`
	Observation string `bson:"observation" json:"observation"`
}
