package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			c.AbortWithStatusJSON(err.Meta.(int), gin.H{
				"error": err.Error(),
			})

		}

	}
}
