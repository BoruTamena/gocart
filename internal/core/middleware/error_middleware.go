package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		// error handling
		if len(c.Errors) > 0 {

			for _, err := range c.Errors {

				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
}
