package usecase

import (
	"api-boilerplate/domain"
	"time"
)

type boilerplateUsecase struct {
	boilerplateRepo domain.BoilerPlateRepository
	contextTimeout  time.Duration
}

func BoilerPlateUsecase(boilerplateRepos domain.BoilerPlateRepository, timeout time.Duration) domain.BoilerPlateUsecase {
	return &boilerplateUsecase{
		boilerplateRepo: boilerplateRepos,
		contextTimeout:  timeout,
	}
}
