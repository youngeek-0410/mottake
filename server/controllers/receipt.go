package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

var (
	ErrInvalidJSONRequest        = errors.New("invalid json request")
	ErrCouldNotCreateReceipt     = errors.New("could not create receipt")
	ErrCouldNotRegisterPurchases = errors.New("could not register purchases")
	ErrCouldNotQueryReceipts     = errors.New("could not query receipts")
)

type ReceiptController struct{}

var receiptModel = new(models.ReceiptModel)

func (r ReceiptController) All(c *gin.Context) {
	uid := getUID(c)
	receipts, err := receiptModel.All(uid)
	if err != nil {
		_ = c.Error(ErrCouldNotQueryReceipts).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, receipts)
}

func (r ReceiptController) Create(c *gin.Context) {
	customerPurchases := struct {
		UID       string            `json:"uid"`
		Purchases []models.Purchase `json:"purchases"`
	}{}
	err := json.NewDecoder(c.Request.Body).Decode(&customerPurchases)
	if err != nil {
		_ = c.Error(ErrInvalidJSONRequest).SetType(gin.ErrorTypePublic)
		return
	}

	receiptID, err := receiptModel.Create(customerPurchases.UID)
	if err != nil {
		_ = c.Error(ErrCouldNotCreateReceipt).SetType(gin.ErrorTypePublic)
		return
	}
	err = receiptModel.RegisterPurchases(receiptID, customerPurchases.Purchases)
	if err != nil {
		_ = c.Error(ErrCouldNotRegisterPurchases).SetType(gin.ErrorTypePublic)
		return
	}
	receipt, _ := receiptModel.GetOne(receiptID, customerPurchases.UID)

	c.JSON(http.StatusCreated, receipt)
}
