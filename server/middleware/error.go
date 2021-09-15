package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/controllers"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		log.Println(err)
		if err != nil {
			apierror := err.Meta.(controllers.APIError)
			c.AbortWithStatusJSON(apierror.StatusCode, gin.H{
				"error": apierror.ErrorMessage,
			})

		}

	}
}
