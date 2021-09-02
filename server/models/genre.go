package models

type Genre struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	FavoriteGenres []FavoriteGenre
	RelatedGenres  []RelatedGenre
}

type FavoriteGenre struct {
	Id          int `json:"id" gorm:"primaryKey"`
	Customer    Customer
	CustomerUid string `json:"customer_uid"`
	Genre       Genre
	GenreId     int `json:"genre_id"`
}

type RelatedGenre struct {
	Id      int `json:"id" gorm:"primaryKey"`
	Shop    Shop
	ShopUid string `json:"shop_Uid"`
	Genre   Genre
	GenreId int `json:"genre_id"`
}
