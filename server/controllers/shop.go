package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/db"
	"github.com/youngeek-0410/mottake/server/models"
)

type ShopController struct{}

func (i ShopController) Get(c *gin.Context) {
	var shop models.Shop
	uid := c.Param("uid")
	if err := db.Db.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Post(c *gin.Context) {
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	shop.UID = getUID(c)
	if err := db.Db.Create(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Patch(c *gin.Context) {
	var shop models.Shop
	uid := getUID(c)
	if err := db.Db.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	shop.UID = uid
	if err := db.Db.Updates(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Delete(c *gin.Context) {
	var shop models.Shop
	uid := getUID(c)
	if err := db.Db.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	shop.UID = uid
	if err := db.Db.Delete(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, shop)
}
