package usecase

import (
	"api-boilerplate/app/helpers"
	"api-boilerplate/app/helpers/validator"
	"api-boilerplate/domain"
	"context"
	"encoding/json"
)

func (boilerplateUC *boilerplateUsecase) ValidationLogin(ctx context.Context, requiredField map[string]string, request domain.RequestLogin, errMessage string) (response domain.Response) {
	var validation map[string]interface{}
	if len(requiredField) > 0 {
		for i, v := range requiredField {
			if len(v) < 1 {
				validation[i] = validator.Required(i)
			}
		}
	}
	if len(request.Email) < 1 {
		validation["email"] = validator.Required("email")
	}
	if len(request.Password) < 1 {
		validation["password"] = validator.Required("password")
	}
	if len(validation) > 0 {
		return helpers.ErrorResponse(400, errMessage, nil, validation)
	}
	return helpers.SuccessResponse("success", json.NewEncoder(nil))
}
