package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	boilerplateHandler "api-boilerplate/boilerplate/delivery/http"
	boilerplateRepo "api-boilerplate/boilerplate/repository"
	boilerplateUseCase "api-boilerplate/boilerplate/usecase"

	"api-boilerplate/app/helpers"
	db "api-boilerplate/app/helpers/services"
	"api-boilerplate/boilerplate/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(*helpers.ProjectFolder + ".env")

	gin.SetMode(gin.ReleaseMode)
	ginC := gin.Default()
	timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	timeoutContext := time.Duration(timeout) * time.Second

	// middleware
	middle := middleware.InitMiddleware()
	ginC.Use(middle.CORS())
	ginC.Use(middle.ValidateContentType())

	// db
	mongoDB, errMongo := db.MongoDB(timeoutContext)
	if errMongo.Status != 200 {
		log.Fatal(errMongo)
	}

	if os.Getenv("GO_ENV") == "development" {
		ginC.Use(gin.Recovery())
	} else if os.Getenv("GO_ENV") == "production" {
		ginC.Use(helpers.CustomRecovery())
	}

	// prefix
	prefixURL := "/boilerplate"

	// boilerplate
	boilerplate := boilerplateRepo.BoilerPlateRepo(mongoDB)
	boilerplateUC := boilerplateUseCase.BoilerPlateUsecase(boilerplate, timeoutContext)
	boilerplateHandler.BoilerPlateHandler(ginC, boilerplateUC, prefixURL)

	port := os.Getenv("PORT")
	fmt.Printf("[%s] API BOILERPLATE running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), port)
	ginC.Run(":" + port)
}
