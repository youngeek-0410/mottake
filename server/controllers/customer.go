package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

type CustomerController struct{}

var customerModel = new(models.CustomerModel)

func (cus CustomerController) Show(c *gin.Context) {
	uid := getUID(c)
	customer, err := customerModel.GetOneByID(uid)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (cus CustomerController) Create(c *gin.Context) {
	var customer models.Customer
	uid := getUID(c)
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}

	reCustomer, err := customerModel.Create(uid, customer)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errCouldNotCreateCustomer})
		return
	}
	c.JSON(http.StatusOK, reCustomer)
}

func (cus CustomerController) Update(c *gin.Context) {
	var customer models.Customer
	uid := getUID(c)
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}

	reCustomer, err := customerModel.Update(uid, customer)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errCouldNotUpdateCustomer})
		return
	}
	c.JSON(http.StatusOK, reCustomer)
}

func (cus CustomerController) Destroy(c *gin.Context) {
	uid := getUID(c)
	if err := customerModel.Destroy(uid); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errCouldNotDeleteCustomer})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
