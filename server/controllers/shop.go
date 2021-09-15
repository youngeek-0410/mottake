package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

var (
	ErrNotFound           = "Not Found"
	ErrCouldNotCreateShop = "Could Not Create Shop"
	ErrInvalidName        = "Invalid character or length (1~30)"
	ErrInvalidDescription = "Invalid character of length (1~255)"
	ErrInvalidAddress     = "Invalid Address. Please enter a proper address in Japanese"
	ErrCouldNotUpdateShop = "Could Not Update Shop"
	ErrInvalidJSON        = "Bad Request (invalid json)"
	ErrNotAuthorized      = "Not Authorized"
	ErrCouldNotDeleteShop = "Could Not Delete Shop"
)

type ShopController struct{}

var shopModel = new(models.ShopModel)

func (i ShopController) GetByID(c *gin.Context) {
	uid := c.Param("uid")
	shop, err := shopModel.GetByID(uid)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, ErrNotFound})
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (i ShopController) Create(c *gin.Context) {
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidJSON})
		return
	}

	// Nameのバリデーション
	if match := NameRegexp.MatchString(shop.Name); match == false {
		c.Error(nil).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidName})
		return
	}
	// Descriptionのバリデーション
	if match := DescriptionRegexp.MatchString(shop.Description); match == false {
		c.Error(nil).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidDescription})
		return
	}
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidAddress})
		return
	}
	uid := getUID(c)
	latitude := coordinate.Latitude
	longitude := coordinate.Longitude

	returnedShop, err := shopModel.Create(shop, uid, latitude, longitude)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, ErrCouldNotCreateShop})
		return
	}
	c.JSON(http.StatusOK, returnedShop)
}

func (i ShopController) Update(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	if err := c.BindJSON(&shop); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidJSON})
		return
	}

	// Nameのバリデーション
	// Nameが空の時はバリデーションを回避（ゼロ値はUpdateされないため）
	if shop.Name != "" && NameRegexp.MatchString(shop.Name) == false {
		c.Error(nil).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidName})
		return
	}
	// Descriptionのバリデーション
	// Nameが空の時はバリデーションを回避（ゼロ値はUpdateされないため）
	if shop.Name != "" && DescriptionRegexp.MatchString(shop.Description) == false {
		c.Error(nil).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidDescription})
		return
	}
	coordinate, err := AddressToCoordinate(shop.Address)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, ErrInvalidAddress})
		return
	}
	latitude := coordinate.Longitude
	longitude := coordinate.Longitude

	returnedshop, err := shopModel.Update(shop, uid, latitude, longitude)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, ErrCouldNotUpdateShop})
		return
	}
	c.JSON(http.StatusOK, returnedshop)
}

func (i ShopController) Delete(c *gin.Context) {
	uid := getUID(c)
	var shop models.Shop
	shop, err := shopModel.Delete(shop, uid)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, ErrCouldNotDeleteShop})
		return
	}
	c.JSON(http.StatusOK, shop)
}
