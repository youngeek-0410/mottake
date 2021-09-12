package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/db"
	"github.com/youngeek-0410/mottake/server/models"
)

type ShopController struct{}

func (i ShopController) Get(c *gin.Context) {
	var shop models.Shop
	uid := c.Param("uid")
	if err := db.DB.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusNotFound, nil)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Post(c *gin.Context) {
	var shop models.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	shop.UID = getUID(c)
	if shop.Coordinate, err = AddressToCoordinate(shop.Address); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	if err := db.DB.Create(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Patch(c *gin.Context) {
	var shop models.Shop
	uid := getUID(c)
	if err := db.DB.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	shop.UID = uid
	if shop.Coordinate, err = AddressToCoordinate(shop.Address); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	if err := db.DB.Updates(&shop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Delete(c *gin.Context) {
	var shop models.Shop
	uid := getUID(c)
	if err := db.DB.Where("UID=?", uid).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	var emptyShop models.Shop
	if err := db.DB.Delete(&emptyShop, uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}
