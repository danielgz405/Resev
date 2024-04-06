package utils

import (
	"context"
	"log"

	"github.com/danielgz405/Resev/database"
	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DatabaseConnection(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DbURI)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	_, err = repository.GetRoleByName(context.Background(), "admin")
	if err != nil && err.Error() == "mongo: no documents in result" {
		id, _ := primitive.ObjectIDFromHex("66108c9f08b2fd8539082911")
		role, err := repository.InsertRole(context.Background(), &models.InsertRole{
			Id:          id,
			Name:        "admin",
			Description: "The admin role",
		})
		if err != nil {
			return
		}
		err = repository.AuditOperation(context.Background(), role.Id.Hex(), "role", "insert")
		if err != nil {
			return
		}
	}

	_, err = repository.GetRoleByName(context.Background(), "client")
	if err != nil && err.Error() == "mongo: no documents in result" {
		id, _ := primitive.ObjectIDFromHex("66108c9f08b2fd8539082913")
		role, err := repository.InsertRole(context.Background(), &models.InsertRole{
			Id:          id,
			Name:        "client",
			Description: "The client role",
		})
		if err != nil {
			return
		}
		err = repository.AuditOperation(context.Background(), role.Id.Hex(), "role", "insert")
		if err != nil {
			return
		}
	}

}

func DatabaseConnection_2(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DB_URI_2)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}

func DatabaseConnection_3(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DB_URI_3)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}
