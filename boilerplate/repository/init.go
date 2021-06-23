package repository

import (
	"api-boilerplate/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type boilerplateRepository struct {
	Conn *mongo.Database
}

func BoilerPlateRepo(Conn *mongo.Database) domain.BoilerPlateRepository {
	return &boilerplateRepository{Conn}
}
