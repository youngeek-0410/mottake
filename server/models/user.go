package models

type User struct {
	Uid  string `json:"uid" gorm:"primaryKey"`
	Name string `json:"name"`
}
