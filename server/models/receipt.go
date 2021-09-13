package models

import (
	"time"

	"github.com/youngeek-0410/mottake/server/db"
	"gorm.io/gorm"
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
	result := db.DB.Create(&receipt)
	if err := result.Error; err != nil {
		return -1, err
	}
	return receipt.ID, nil
}

func (r ReceiptModel) RegisterPurchases(receiptID int, purchases []Purchase) error {
	for _, purchase := range purchases {
		purchase.ReceiptID = receiptID
		result := db.DB.Create(&purchase)
		if err := result.Error; err != nil {
			return err
		}
	}
	return nil
}

func (r ReceiptModel) GetOne(receiptID int, customerUID string) (*Receipt, error) {
	var receipt Receipt
	var result *gorm.DB

	result = db.DB.Where("id = ? AND customer_uid = ?", receiptID, customerUID).Preload("Purchases").First(&receipt)

	if err := result.Error; err != nil {
		return nil, err
	}
	return &receipt, nil
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
