package usecase

import (
	"api-boilerplate/domain"
	"context"
)

func (boilerplateUC *boilerplateUsecase) OptionStatus(ctx context.Context) (response []string) {
	return []string{"publish", "draft", "archive", "trash"}
}

func (boilerplateUC *boilerplateUsecase) ConvertUserResp(ctx context.Context, mongo domain.User) (response domain.User) {
	return response
}

func (boilerplateUC *boilerplateUsecase) ConvertUserRespMultiple(ctx context.Context, mongo []domain.User) (response []domain.User) {
	response = make([]domain.User, 0)
	for _, v := range mongo {
		item := v
		response = append(response, item)
	}
	return response
}
