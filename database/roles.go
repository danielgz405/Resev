package database

import (
	"context"

	"github.com/danielgz405/Resev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertRole(ctx context.Context, role *models.InsertRole) (*models.Role, error) {
	collection := repo.client.Database("resev").Collection("role")
	result, err := collection.InsertOne(ctx, role)
	if err != nil {
		return nil, err
	}
	createdRole, err := repo.GetRoleById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdRole, nil
}

func (repo *MongoRepo) GetRoleById(ctx context.Context, id string) (*models.Role, error) {
	collection := repo.client.Database("resev").Collection("role")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (repo *MongoRepo) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
	collection := repo.client.Database("resev").Collection("role")
	var role models.Role
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (repo *MongoRepo) ListRoles(ctx context.Context) ([]models.Role, error) {
	collection := repo.client.Database("resev").Collection("role")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	roles := []models.Role{}

	if err = cursor.All(ctx, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (repo *MongoRepo) UpdateRole(ctx context.Context, data *models.UpdateRole, id string) (*models.Role, error) {
	collection := repo.client.Database("resev").Collection("role")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":        data.Name,
		"description": data.Description,
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
	updatedRole, err := repo.GetRoleById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedRole, nil
}

func (repo *MongoRepo) DeleteRole(ctx context.Context, id string) error {
	collection := repo.client.Database("resev").Collection("role")
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

func (repo *MongoRepo) ListRolesByPage(ctx context.Context, limit int, page int) ([]models.Role, int, error) {
	collection := repo.client.Database("resev").Collection("role")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	roles := []models.Role{}
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, 0, err
	}

	return roles, int(quantity), nil
}
