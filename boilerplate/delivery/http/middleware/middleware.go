package middleware

import (
	"api-boilerplate/app/helpers"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GoMiddleware struct {
}

func (m *GoMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func (m *GoMiddleware) ValidateContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !helpers.InArrayString(c.Request.Method, []string{"GET", "DELETE"}) {
			contentType := c.Request.Header.Get("Content-Type")
			if contentType != "application/json" {
				errMessage := "This service only support for application/json content type"
				response := helpers.ErrorResponse(400, errMessage, errors.New(errMessage), map[string]interface{}{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			}
		}
		c.Next()
	}
}

func (m *GoMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hAuth := c.GetHeader("Authorization")
		if len(hAuth) < 1 {
			errMessage := "Unauthorized"
			response := helpers.ErrorResponse(400, errMessage, errors.New(errMessage), map[string]interface{}{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		splitToken := strings.Split(hAuth, "Bearer ")
		if len(splitToken) != 2 {
			errMessage := "Unauthorized"
			response := helpers.ErrorResponse(400, errMessage, errors.New(errMessage), map[string]interface{}{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		tokenString := splitToken[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if !token.Valid {
			errMessage := err.Error()
			response := helpers.ErrorResponse(400, errMessage, errors.New(errMessage), map[string]interface{}{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		c.Set("user", claims["user"])
		c.Next()
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
