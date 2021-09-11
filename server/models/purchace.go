package models

type Purchase struct {
	ID        int      `json:"id" gorm:"primaryKey"`
	MenuID    int      `json:"menu_id"`
	Receipt   *Receipt `json:"receipt,omitempty"`
	ReceiptID int      `json:"receipt_id"`
	Number    int      `json:"number"`
}
