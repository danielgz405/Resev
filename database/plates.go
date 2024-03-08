package database

import (
	"context"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertPlate(ctx context.Context, plate *models.InsertPlate) (*models.Plate, error) {
	collection := repo.client.Database("resev").Collection("plate")
	result, err := collection.InsertOne(ctx, plate)
	if err != nil {
		return nil, err
	}
	createdPlate, err := repo.GetPlateById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdPlate, nil
}

func (repo *MongoRepo) GetPlateById(ctx context.Context, id string) (*models.Plate, error) {
	collection := repo.client.Database("resev").Collection("plate")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var plate models.Plate
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&plate)
	if err != nil {
		return nil, err
	}
	return &plate, nil
}

func (repo *MongoRepo) ListPlates(ctx context.Context) ([]models.Plate, error) {
	collection := repo.client.Database("resev").Collection("plate")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	plates := []models.Plate{}

	if err = cursor.All(ctx, &plates); err != nil {
		return nil, err
	}
	return plates, nil
}

func (repo *MongoRepo) UpdatePlate(ctx context.Context, data *models.UpdatePlate, id string) (*models.Plate, error) {
	collection := repo.client.Database("resev").Collection("plate")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":      data.Name,
		"price":     data.Price,
		"details":   data.Details,
		"image":     data.Image,
		"image_ref": data.ImageRef,
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
	updatedPlate, err := repo.GetPlateById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedPlate, nil
}

func (repo *MongoRepo) DeletePlate(ctx context.Context, id string) error {
	collection := repo.client.Database("resev").Collection("plate")
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

func (repo *MongoRepo) ListPlatesByPage(ctx context.Context, limit int, page int) ([]models.Plate, int, error) {
	collection := repo.client.Database("resev").Collection("plate")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	plates := []models.Plate{}
	if err = cursor.All(ctx, &plates); err != nil {
		return nil, 0, err
	}

	return plates, int(quantity), nil
}

func (repo *MongoRepo) GetPlatesByIds(ctx context.Context, ids []string) ([]models.Plate, error) {
	collection := repo.client.Database("resev").Collection("plate")

	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, oid)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	plates := []models.Plate{}
	if err = cursor.All(ctx, &plates); err != nil {
		return nil, err
	}

	return plates, nil
}
