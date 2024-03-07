package database

import (
	"context"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertTable(ctx context.Context, table *models.InsertTable) (*models.Table, error) {
	collection := repo.client.Database("resev").Collection("table")
	result, err := collection.InsertOne(ctx, table)
	if err != nil {
		return nil, err
	}
	createdTable, err := repo.GetTableById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdTable, nil
}

func (repo *MongoRepo) GetTableById(ctx context.Context, id string) (*models.Table, error) {
	collection := repo.client.Database("resev").Collection("table")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var table models.Table
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (repo *MongoRepo) ListTables(ctx context.Context) ([]models.Table, error) {
	collection := repo.client.Database("resev").Collection("table")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	tables := []models.Table{}

	if err = cursor.All(ctx, &tables); err != nil {
		return nil, err
	}
	return tables, nil
}

func (repo *MongoRepo) UpdateTable(ctx context.Context, data *models.UpdateTable, id string) (*models.Table, error) {
	collection := repo.client.Database("resev").Collection("table")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"number":      data.Number,
		"observation": data.Observation,
	}
	for key, value := range iterableData {
		if value != nil && value != "" {
			update["$set"].(bson.M)[key] = value
		}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	updatedTable, err := repo.GetTableById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedTable, nil
}

func (repo *MongoRepo) DeleteTable(ctx context.Context, id string) error {
	collection := repo.client.Database("resev").Collection("table")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoRepo) ListTablesByPage(ctx context.Context, limit int, page int) ([]models.Table, int, error) {
	collection := repo.client.Database("resev").Collection("table")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	tables := []models.Table{}
	if err = cursor.All(ctx, &tables); err != nil {
		return nil, 0, err
	}

	return tables, int(quantity), nil
}
