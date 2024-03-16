package utils

import (
	"log"

	"github.com/danielgz405/Resev/database"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/server"
)

func DatabaseConnection(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DbURI)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}

func DatabaseConnection_2(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DB_URI_2)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}
