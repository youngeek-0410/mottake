package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

type ReceiptController struct{}

var receiptModel = new(models.ReceiptModel)

func (r ReceiptController) All(c *gin.Context) {
	uid := getUID(c)
	receipts, err := receiptModel.All(uid)
	if err != nil {
		_ = c.Error(errCouldNotQueryReceipts).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, receipts)
}

func (r ReceiptController) Create(c *gin.Context) {
	var receipt models.Receipt
	shopUID := getUID(c)

	if err := c.ShouldBindJSON(&receipt); err != nil {
		_ = c.Error(errInvalidJSONRequest).SetType(gin.ErrorTypePublic)
		return
	}

	receiptID, err := receiptModel.Create(receipt, shopUID)
	if err != nil {
		_ = c.Error(errCouldNotCreateReceipt).SetType(gin.ErrorTypePublic)
		return
	}
	receipt, _ = receiptModel.GetOneByID(receiptID, receipt.CustomerUID)

	c.JSON(http.StatusCreated, receipt)
}
