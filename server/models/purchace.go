package models

type Purchase struct {
	Id     int `json:"id" gorm:"primaryKey"`
	MenuId int `json:"menu_id"`
	Number int `json:"number"`
}
