package models

import (
	"github.com/youngeek-0410/mottake/server/db"
	"gorm.io/gorm"
	"time"
)

type Receipt struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerUID string     `json:"customer_uid"`
	CreatedAt   time.Time  `json:"created_at"`
	Purchases   []Purchase `json:"purchases,omitempty"`
}

type ReceiptModel struct{}

func (r ReceiptModel) Create(customerUID string) (receiptID int, err error) {
	receipt := Receipt{
		CustomerUID: customerUID,
		CreatedAt:   time.Now(),
	}
	result := db.Db.Create(&receipt)
	if err := result.Error; err != nil {
		return -1, err
	}
	return receipt.ID, nil
}

func (r ReceiptModel) RegisterPurchases(receiptID int, purchases []Purchase) error {
	for _, purchase := range purchases {
		purchase.ReceiptID = receiptID
		result := db.Db.Create(&purchase)
		if err := result.Error; err != nil {
			return err
		}
	}
	return nil
}

func (r ReceiptModel) GetOneByID(receiptID int, preload bool) (*Receipt, error) {
	var receipt Receipt
	var result *gorm.DB
	if preload {
		result = db.Db.Where("id = ?", receiptID).Preload("Purchases").First(&receipt)
	} else {
		result = db.Db.Where("id = ?", receiptID).First(&receipt)
	}

	if err := result.Error; err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (r ReceiptModel) All(customerUID string, preload bool) ([]Receipt, error) {
	var receipts []Receipt
	var result *gorm.DB
	if preload {
		result = db.Db.Where("customer_uid = ?", customerUID).Preload("Purchases").Find(&receipts)
	} else {
		result = db.Db.Where("customer_uid = ?", customerUID).Find(&receipts)
	}
	if err := result.Error; err != nil {
		return nil, err
	}
	return receipts, nil
}
