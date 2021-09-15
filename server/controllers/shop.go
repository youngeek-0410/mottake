package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

var (
	ErrNotFound           = errors.New("Not Found")
	ErrCouldNotCreateShop = errors.New("Could Not Create Shop")
	ErrInvalidName        = errors.New("Invalid character or length (1~30)")
	ErrInvalidDescription = errors.New("Invalid character of length (1~255)")
	ErrInvalidAddress     = errors.New("Invalid Address. Please enter a proper address in Japanese")
	ErrCouldNotUpdateShop = errors.New("Could Not Update Shop")
	ErrInvalidJSON        = errors.New("Bad Request (invalid json)")
	ErrNotAuthorized      = errors.New("Not Authorized")
	ErrCouldNotDeleteShop = errors.New("Could Not Delete Shop")
)

type ShopController struct{}

var shopModel = new(models.ShopModel)

func (i ShopController) Get(c *gin.Context) {
	uid := c.Param("uid")
	shop, err := shopModel.GetByID(uid)
	if err != nil {
		log.Println(err)
		c.Error(ErrNotFound).SetType(gin.ErrorTypePublic).SetMeta(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Post(c *gin.Context) {
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		log.Println(err)
		c.Error(ErrInvalidJSON).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	// Nameのバリデーション
	if match := NameRegexp.MatchString(shop.Name); match == false {
		c.Error(ErrInvalidName).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	// Descriptionのバリデーション
	if match := DescriptionRegexp.MatchString(shop.Description); match == false {
		c.Error(ErrInvalidDescription).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		log.Println(err)
		c.Error(ErrInvalidAddress).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	uid := getUID(c)
	latitude := coordinate.Latitude
	longitude := coordinate.Longitude

	returnedShop, err := shopModel.Create(shop, uid, latitude, longitude)
	if err != nil {
		log.Println(err)
		c.Error(ErrCouldNotCreateShop).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, returnedShop)
}

func (i ShopController) Patch(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		log.Println(err)
		c.Error(ErrInvalidJSON).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	// Nameのバリデーション
	if match := NameRegexp.MatchString(shop.Name); match != false {
		c.Error(ErrInvalidName).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	// Descriptionのバリデーション
	if match := DescriptionRegexp.MatchString(shop.Description); match != false {
		c.Error(ErrInvalidDescription).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	//if shop.Address != "" {
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		log.Println(err)
		c.Error(ErrInvalidAddress).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	latitude := coordinate.Longitude
	longitude := coordinate.Longitude
	//}

	shop, err = shopModel.Update(shop, uid, latitude, longitude)
	if err != nil {
		log.Println(err)
		c.Error(ErrCouldNotUpdateShop).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Delete(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	shop, err := shopModel.Delete(shop, uid)
	if err != nil {
		log.Println(err)
		c.Error(ErrCouldNotDeleteShop).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, shop)
}
