package models

import (
	"github.com/youngeek-0410/mottake/server/db"
)

type Genre struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	//	FavoriteGenres []FavoriteGenre
	//	RelatedGenres []RelatedGenre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type FavoriteGenre struct {
	CustomerUID string `json:"customer_uid" gorm:"primaryKey"`
	Genre       Genre  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GenreID     int    `json:"genre_id"`
}

type RelatedGenre struct {
	ShopUID string `json:"shop_uid" gorm:"primaryKey"`
	Genre   Genre  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GenreID int    `json:"genre_id"`
}

type FavoriteGenreModel struct{}

func (r FavoriteGenreModel) GetByID(uid string) (FavoriteGenre, error) {
	var genre FavoriteGenre
	if err := db.DB.Where("customer_uid=?", uid).Preload("Genre").First(&genre).Error; err != nil {
		return genre, err
	}
	return genre, nil
}

func (r FavoriteGenreModel) Save(genre FavoriteGenre) (FavoriteGenre, error) {
	if err := db.DB.Save(&genre).Error; err != nil {
		return genre, err
	}
	return genre, nil
}

type RelatedGenreModel struct{}

func (r RelatedGenreModel) GetByID(uid string) (RelatedGenre, error) {
	var genre RelatedGenre
	if err := db.DB.Where("shop_uid=?", uid).Preload("Genre").First(&genre).Error; err != nil {
		return genre, err
	}
	return genre, nil
}

func (r RelatedGenreModel) Save(genre RelatedGenre) (RelatedGenre, error) {
	if err := db.DB.Save(&genre).Error; err != nil {
		return genre, err
	}
	return genre, nil
}
