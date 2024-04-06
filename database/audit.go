package database

import (
	"context"
	"time"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) AuditOperation(ctx context.Context, id string, table string, operationType string) error {
	//get document
	collection := repo.client.Database("resev").Collection(table)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var document interface{}
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&document)
	if err != nil {
		return err
	}

	auditLog := models.AuditLog{
		Id:            primitive.NewObjectID(),
		Document:      document,
		OperationType: operationType,
		OperationDate: time.Now().UTC().Format("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)"),
		ClusterTime: primitive.Timestamp{
			T: uint32(time.Now().Unix()),
			I: 0,
		},

		WallTime: time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
	}

	collection = repo.client.Database("resev").Collection(table + "_audit")
	_, err = collection.InsertOne(ctx, auditLog)
	if err != nil {
		return err
	}

	return nil
}
