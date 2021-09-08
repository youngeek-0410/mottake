package models

type Purchase struct {
	ID     int `json:"id" gorm:"primaryKey"`
	MenuID int `json:"menu_id"`
	Number int `json:"number"`
}
