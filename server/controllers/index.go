package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (i IndexController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
