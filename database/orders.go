package database

import (
	"context"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertOrder(ctx context.Context, order *models.InsertOrder) (*models.Order, error) {
	collection := repo.client.Database("resev").Collection("order")
	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	createdOrder, err := repo.GetOrderById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdOrder, nil
}

func (repo *MongoRepo) GetOrderById(ctx context.Context, id string) (*models.Order, error) {
	collection := repo.client.Database("resev").Collection("order")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var order models.Order
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *MongoRepo) ListOrders(ctx context.Context) ([]models.Order, error) {
	collection := repo.client.Database("resev").Collection("order")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	orders := []models.Order{}

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *MongoRepo) UpdateOrder(ctx context.Context, data *models.UpdateOrder, id string) (*models.Order, error) {
	collection := repo.client.Database("resev").Collection("order")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"table_id":  data.TableId,
		"subtotal":  data.SubTotal,
		"total":     data.Total,
		"timeStamp": data.TimeStamp,
		"plates_id": data.PlatesId,
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
	updatedOrder, err := repo.GetOrderById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}

func (repo *MongoRepo) DeleteOrder(ctx context.Context, id string) error {
	collection := repo.client.Database("resev").Collection("order")
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

func (repo *MongoRepo) ListOrdersByPage(ctx context.Context, limit int, page int) ([]models.Order, int, error) {
	collection := repo.client.Database("resev").Collection("order")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	orders := []models.Order{}
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, 0, err
	}

	return orders, int(quantity), nil
}
