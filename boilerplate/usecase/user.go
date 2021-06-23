package usecase

import (
	"api-boilerplate/app/helpers"
	"api-boilerplate/app/helpers/validator"
	"api-boilerplate/domain"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (boilerplateUC *boilerplateUsecase) Login(c context.Context, request domain.RequestLogin) (response domain.Response) {
	ctx, cancel := context.WithTimeout(c, boilerplateUC.contextTimeout)
	defer cancel()

	// VARIABLE
	var jwtISS string
	var jwtSecretKey string
	validation := make(map[string]interface{})
	errMessage := "Sorry, Email or password is invalid"

	// VALIDATION
	check := boilerplateUC.ValidationLogin(ctx, map[string]string{}, request, errMessage)
	if check.Status != 200 {
		return check
	}

	// DO LOGIN
	options := map[string]interface{}{
		"email":    request.Email,
		"password": request.Password,
	}
	data, err := boilerplateUC.boilerplateRepo.Login(ctx, options)
	if err != nil {
		return helpers.ErrorResponse(400, errMessage, err, validation)
	}

	// CHECK PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("ERROR (CompareHashAndPassword): ", err)
		return helpers.ErrorResponse(400, errMessage, nil, validation)
	}

	// JWT
	jwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	jwtISS = os.Getenv("JWT_ISS")
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	timeNow := time.Now()
	claims := sign.Claims.(jwt.MapClaims)

	claims["iat"] = timeNow.Unix()
	claims["iss"] = jwtISS
	claims["aud"] = ""
	claims["sub"] = request.Email

	user := make(map[string]interface{})
	user["id"] = data.ID
	claims["user"] = user
	jwtToken, _ := sign.SignedString([]byte(jwtSecretKey))
	user["token"] = jwtToken
	return helpers.SuccessResponse("success", user)
}

func (boilerplateUC *boilerplateUsecase) Detail(c context.Context, ID string) domain.Response {
	ctx, cancel := context.WithTimeout(c, boilerplateUC.contextTimeout)
	defer cancel()

	// VARIABLE
	validation := make(map[string]interface{})
	errMessage := validator.FailedToGetData("detail")

	// VALIDATION
	if len(ID) < 1 {
		validation["ID"] = validator.Required("ID")
	}
	if len(validation) > 0 {
		return helpers.ErrorResponse(400, errMessage, nil, validation)
	}
	ids, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		fmt.Println("ERROR (ObjectIDFromHex ID): ", ID)
		return helpers.ErrorResponse(400, errMessage, err, validation)
	}

	// GET DETAIL USER
	data, err := boilerplateUC.boilerplateRepo.Detail(ctx, ids)
	if err != nil {
		return helpers.ErrorResponse(400, errMessage, err, validation)
	}

	// CONVERT RESPONSE
	response := boilerplateUC.ConvertUserResp(ctx, data)
	return helpers.SuccessResponse("success", response)
}
