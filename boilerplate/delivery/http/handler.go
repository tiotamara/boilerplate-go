package http

import (
	helpers "api-boilerplate/app/helpers"
	"api-boilerplate/boilerplate/delivery/http/middleware"
	"api-boilerplate/domain"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type boilerplateHandler struct {
	BoilerPlateUsecase domain.BoilerPlateUsecase
}

func BoilerPlateHandler(g *gin.Engine, boilerplate domain.BoilerPlateUsecase, prefix string) {
	middle := middleware.InitMiddleware()
	handler := &boilerplateHandler{
		BoilerPlateUsecase: boilerplate,
	}
	v1 := g.Group(prefix + "/v1/boilerplate")
	v1.GET("/", handler.Index)
	v1.POST("/login", handler.Login)

	// program
	program := v1.Group("/users")
	program.GET("/detail/:id", middle.Auth(), handler.Detail)
}

func (a *boilerplateHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.SuccessResponse("API BOILERPLATE :)", json.NewEncoder(nil)))
	return
}

func (a *boilerplateHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	errMessage := "Sorry email or password not valid"
	var request domain.RequestLogin

	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse(http.StatusBadRequest, errMessage, err, map[string]interface{}{}))
		return
	}
	response := a.BoilerPlateUsecase.Login(ctx, request)
	c.JSON(response.Status, response)
	return
}

func (a *boilerplateHandler) Detail(c *gin.Context) {
	ctx := c.Request.Context()

	user := c.MustGet("user")
	dataJWT := domain.Jwt{}
	byteUser, _ := json.Marshal(user)
	json.Unmarshal(byteUser, &dataJWT)

	response := a.BoilerPlateUsecase.Detail(ctx, dataJWT.ID)
	c.AbortWithStatusJSON(response.Status, response)
	return
}
