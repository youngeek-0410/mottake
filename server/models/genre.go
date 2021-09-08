package models

type Genre struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	FavoriteGenres []FavoriteGenre
	RelatedGenres  []RelatedGenre
}

type FavoriteGenre struct {
	ID          int `json:"id" gorm:"primaryKey"`
	Customer    Customer
	CustomerUID string `json:"customer_uid"`
	Genre       Genre
	GenreID     int `json:"genre_id"`
}

type RelatedGenre struct {
	ID      int `json:"id" gorm:"primaryKey"`
	Shop    Shop
	ShopUID string `json:"shop_Uid"`
	Genre   Genre
	GenreID int `json:"genre_id"`
}
