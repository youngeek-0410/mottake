package models

type Customer struct {
	UID           string        `json:"uid" gorm:"primaryKey"`
	FavoriteGenre FavoriteGenre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
