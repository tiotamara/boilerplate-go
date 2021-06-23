package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BoilerPlateUsecase interface {
	Login(ctx context.Context, payload RequestLogin) Response
	Detail(ctx context.Context, ID string) Response
}

type BoilerPlateRepository interface {
	Login(ctx context.Context, optionFilter map[string]interface{}) (User, error)
	Detail(ctx context.Context, ID primitive.ObjectID) (User, error)
}
