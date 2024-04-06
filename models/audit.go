package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuditLog struct {
	Id            primitive.ObjectID  `bson:"_id" json:"_id"`
	Document      interface{}         `bson:"document" json:"document"`
	OperationType string              `bson:"operation_type" json:"operation_type"`
	OperationDate string              `bson:"operation_date" json:"operation_date"`
	ClusterTime   primitive.Timestamp `bson:"cluster_time" json:"cluster_time"`
	WallTime      string              `bson:"wall_time" json:"wall_time"`
}

type WallTime struct {
	Date string `bson:"date" json:"date"`
}
