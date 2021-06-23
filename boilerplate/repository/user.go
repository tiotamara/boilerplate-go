package repository

import (
	"api-boilerplate/app/helpers"
	"api-boilerplate/domain"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateQueryFilterUser(options map[string]interface{}) bson.M {
	query := bson.M{
		"deleted_at": bson.M{
			"$eq": nil,
		},
	}

	if id, ok := options["id"].(primitive.ObjectID); ok {
		query["_id"] = id
	}
	if email, ok := options["email"].(string); ok {
		query["email"] = email
	}

	return query
}

func (m *boilerplateRepository) Login(ctx context.Context, options map[string]interface{}) (response domain.User, err error) {
	query := generateQueryFilterUser(options)

	collection := m.Conn.Collection(os.Getenv("MONGO_COLLECTION_USER"))
	err = collection.FindOne(context.TODO(), query).Decode(&response)
	if err != nil {
		fmt.Println("ERROR (Login): ", err, query, helpers.PrettyPrint(options))
		return response, err
	}
	return response, err
}

func (m *boilerplateRepository) Detail(ctx context.Context, ID primitive.ObjectID) (response domain.User, err error) {
	options := map[string]interface{}{
		"id": ID,
	}
	query := generateQueryFilterUser(options)

	collection := m.Conn.Collection(os.Getenv("MONGO_COLLECTION_USER"))
	err = collection.FindOne(context.TODO(), query).Decode(&response)
	if err != nil {
		fmt.Println("ERROR (Detail): ", err, query)
		return response, err
	}
	return response, err
}
