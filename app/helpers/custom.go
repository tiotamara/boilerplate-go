package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDateUTC(format string) (time.Time, string) {
	dt := time.Now().UTC()
	dateNow := dt.Format(format)
	return dt, dateNow
}

func GetDateWIB(format string) (time.Time, string) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(location)
	dateNow := dt.Format(format)
	return dt, dateNow
}

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic Recover : ", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse(http.StatusInternalServerError, "Something went wrong", nil, map[string]interface{}{}))
			}
		}()
		c.Next()
	}
}
