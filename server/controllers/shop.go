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
		c.JSON(http.StatusNotFound, ErrNotFound)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Post(c *gin.Context) {
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, ErrInvalidJSON)
		log.Println(err)
		return
	}

	// Nameのバリデーション
	if match := NameRegexp.MatchString(shop.Name); match == false {
		c.JSON(http.StatusBadRequest, ErrInvalidName)
		return
	}
	// Descriptionのバリデーション
	if match := DescriptionRegexp.MatchString(shop.Description); match == false {
		c.JSON(http.StatusBadRequest, ErrInvalidDescription)
		return
	}
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrInvalidAddress)
		log.Println(err)
		return
	}
	uid := getUID(c)
	latitude := coordinate.Latitude
	longitude := coordinate.Longitude

	returnedShop, err := shopModel.Create(shop, uid, latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrCouldNotCreateShop)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, returnedShop)
}

func (i ShopController) Patch(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, ErrInvalidJSON)
		log.Println(err)
		return
	}

	// Nameのバリデーション
	if match := NameRegexp.MatchString(shop.Name); match != false {
		c.JSON(http.StatusBadRequest, ErrInvalidName)
		return
	}
	// Descriptionのバリデーション
	if match := DescriptionRegexp.MatchString(shop.Description); match != false {
		c.JSON(http.StatusBadRequest, ErrInvalidDescription)
		return
	}
	//if shop.Address != "" {
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrInvalidAddress)
		log.Println(err)
		return
	}
	latitude := coordinate.Longitude
	longitude := coordinate.Longitude
	//}

	shop, err = shopModel.Update(shop, uid, latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrCouldNotUpdateShop)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Delete(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	shop, err := shopModel.Delete(shop, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrCouldNotDeleteShop)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, shop)
}
