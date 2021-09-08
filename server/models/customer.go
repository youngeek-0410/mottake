package models

type Customer struct {
	Uid            string `json:"uid" gorm:"primaryKey"`
	FavoriteGenres []FavoriteGenre
	Receipts       []Receipt
}
