package models

import (
	"time"

	"github.com/youngeek-0410/mottake/server/db"
	"gorm.io/gorm"
)

type Receipt struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerUID string     `json:"customer_uid" binding:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	Purchases   []Purchase `json:"purchases,omitempty" binding:"required"`
}

type ReceiptModel struct{}

func (r ReceiptModel) Create(receipt Receipt, shopUID string) (receiptID int, err error) {

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		var shop Shop
		var total int
		if err = db.DB.Where("uid = ?", shopUID).First(&shop).Error; err != nil {
			return err
		}
		sales := shop.Sales
		r := Receipt{
			CustomerUID: receipt.CustomerUID,
			CreatedAt:   time.Now(),
		}
		if err := db.DB.Create(&r).Error; err != nil {
			return err
		}
		receiptID = r.ID
		for _, purchase := range receipt.Purchases {
			purchase.ReceiptID = receiptID
			if err := db.DB.Create(&purchase).Error; err != nil {
				return err
			}
			if err = db.DB.Where("id = ?", purchase.ID).Preload("Menu").First(&purchase).Error; err != nil{
				return err
			}
			total = total + (purchase.Menu.Price * purchase.Number)
		}
		shop.Sales = sales + total
		if err = db.DB.Updates(&shop).Error; err != nil {
			return err
		}

		return nil
	})
	return receiptID, err
}

func (r ReceiptModel) GetOneByID(receiptID int, customerUID string) (Receipt, error) {
	var receipt Receipt
	var result *gorm.DB

	result = db.DB.Where("id = ? AND customer_uid = ?", receiptID, customerUID).Preload("Purchases").First(&receipt)

	if err := result.Error; err != nil {
		return receipt, err
	}
	return receipt, nil
}

func (r ReceiptModel) All(customerUID string) ([]Receipt, error) {
	var receipts []Receipt
	var result *gorm.DB
	result = db.DB.Where("customer_uid = ?", customerUID).Preload("Purchases").Find(&receipts)

	if err := result.Error; err != nil {
		return nil, err
	}
	return receipts, nil
}
