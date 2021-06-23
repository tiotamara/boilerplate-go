package services

import (
	"api-boilerplate/app/helpers"
	"api-boilerplate/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoDB(timeout time.Duration) (*mongo.Database, domain.Response) {
	errMessage := "Maaf, Terjadi kesalahaan pada database"
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbAuth := os.Getenv("MONGO_DB_AUTH")
	scheme := "mongodb://"
	if os.Getenv("MONGO_ATLAS") == "true" {
		scheme = "mongodb+srv://"
	}
	userpass := ""
	if user != "" {
		userpass = user + "@"
		if pass != "" {
			userpass = user + ":" + pass + "@"
		}
	}
	if port != "" {
		port = ":" + port
	}
	url := scheme + userpass + host + port + "/" + dbAuth
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	clientOptions := options.Client()
	clientOptions.ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println("MongoDB", err)
		return nil, helpers.ErrorResponse(500, errMessage, errors.New(errMessage), map[string]interface{}{})
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("MongoDB", err)
		return nil, helpers.ErrorResponse(500, errMessage, errors.New(errMessage), map[string]interface{}{})
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("MongoDB", err)
		return nil, helpers.ErrorResponse(500, errMessage, errors.New(errMessage), map[string]interface{}{})
	}
	return client.Database(os.Getenv("MONGO_DB_NAME")), helpers.SuccessResponse("success", json.NewEncoder(nil))
}
