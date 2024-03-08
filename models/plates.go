package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plate struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Price    float64            `bson:"price" json:"price"`
	Details  string             `bson:"details" json:"details"`
	Image    string             `bson:"image" json:"image"`
	ImageRef string             `bson:"image_ref" json:"image_ref"`
}

type InsertPlate struct {
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Details  string  `bson:"details" json:"details"`
	Image    string  `bson:"image" json:"image"`
	ImageRef string  `bson:"image_ref" json:"image_ref"`
}

type UpdatePlate struct {
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Details  string  `bson:"details" json:"details"`
	Image    string  `bson:"image" json:"image"`
	ImageRef string  `bson:"image_ref" json:"image_ref"`
}
