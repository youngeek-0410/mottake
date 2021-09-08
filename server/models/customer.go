package models

type Customer struct {
	UID            string `json:"uid" gorm:"primaryKey"`
	FavoriteGenres []FavoriteGenre
	Receipts       []Receipt
}
