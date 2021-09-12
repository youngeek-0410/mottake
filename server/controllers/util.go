package controllers

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func getUID(c *gin.Context) string {
	token, _ := c.Get("token")
	uid := token.(*auth.Token).UID
	return uid
}
