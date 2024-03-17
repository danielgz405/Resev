package database

import (
	"context"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertBooking(ctx context.Context, booking *models.InsertBooking) (*models.Booking, error) {
	collection := repo.client.Database("resev").Collection("booking")
	result, err := collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	createdBooking, err := repo.GetBookingById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdBooking, nil
}

func (repo *MongoRepo) GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	collection := repo.client.Database("resev").Collection("booking")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking models.Booking
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (repo *MongoRepo) ListBookings(ctx context.Context) ([]models.Booking, error) {
	collection := repo.client.Database("resev").Collection("booking")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	bookings := []models.Booking{}

	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (repo *MongoRepo) UpdateBooking(ctx context.Context, data *models.UpdateBooking, id string) (*models.Booking, error) {
	collection := repo.client.Database("resev").Collection("booking")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"table_id":    data.TableId,
		"name":        data.Name,
		"description": data.Description,
		"date":        data.Date,
		"hour":        data.Hour,
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
	updatedBooking, err := repo.GetBookingById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedBooking, nil
}

func (repo *MongoRepo) DeleteBooking(ctx context.Context, id string) error {
	collection := repo.client.Database("resev").Collection("booking")
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

func (repo *MongoRepo) ListBookingsByPage(ctx context.Context, limit int, page int) ([]models.Booking, int, error) {
	collection := repo.client.Database("resev").Collection("booking")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	bookings := []models.Booking{}
	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, 0, err
	}

	return bookings, int(quantity), nil
}

func (repo *MongoRepo) GetBookingsByIds(ctx context.Context, ids []string) ([]models.Booking, error) {
	collection := repo.client.Database("resev").Collection("booking")

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

	bookings := []models.Booking{}
	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}
