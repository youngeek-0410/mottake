package utils

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) *auth.Token {
	uid, _ := c.Get("token")
	return uid.(*auth.Token)
}
