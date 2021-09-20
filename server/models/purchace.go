package models

type Purchase struct {
	ID        int  `json:"id" gorm:"primaryKey"`
	Menu      Menu `json:"menu,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MenuID    int  `json:"menu_id"`
	ReceiptID int  `json:"receipt_id"`
	Number    int  `json:"number"`
}
