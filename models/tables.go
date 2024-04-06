package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Table struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Number      string             `bson:"number" json:"number"`
	Observation string             `bson:"observation" json:"observation"`
	ImageBase64 primitive.Binary   `bson:"image_base64" json:"image_base64"`
}

type InsertTable struct {
	Number      string           `bson:"number" json:"number"`
	Observation string           `bson:"observation" json:"observation"`
	ImageBase64 primitive.Binary `bson:"image_base64" json:"image_base64"`
}

type UpdateTable struct {
	Number      string           `bson:"number" json:"number"`
	Observation string           `bson:"observation" json:"observation"`
	ImageBase64 primitive.Binary `bson:"image_base64" json:"image_base64"`
}
