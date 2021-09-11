package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShopController struct{}

func (i ShopController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Shop!",
	})
}

func (i ShopController) Post(c *gin.Context) {
	var s user.Service
}
