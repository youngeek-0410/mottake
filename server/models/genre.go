package models

type Genre struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	FavoriteGenres []FavoriteGenre
	RelatedGenres  []RelatedGenre
}

type FavoriteGenre struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	CustomerUid string `json:"customer_uid"`
	GenreId     int    `json:"genre_id"`
}

type RelatedGenre struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	ShopUid string `json:"shop_Uid"`
	GenreId int    `json:"genre_id"`
}
