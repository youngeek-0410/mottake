package models

import (
	"github.com/youngeek-0410/mottake/server/db"
)

type Shop struct {
	UID           string         `json:"uid" gorm:"primaryKey"`
	Name          string         `json:"name"`
	Image         string         `json:"image"`
	Address       string         `json:"address"`
	Latitude      float32        `json:"latitude"`
	Longitude     float32        `json:"longitude"`
	SalesGoal     int            `json:"sales_goal"`
	Sales         int            `json:"sales"`
	Description   string         `json:"description"`
	Menus         []Menu         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RelatedGenres []RelatedGenre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ShopModel struct{}

func (r ShopModel) GetByID(uid string) (Shop, error) {
	var shop Shop
	if err := db.DB.Where("UID=?", uid).Preload("RelatedGenres").First(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func (r ShopModel) Create(shop Shop, uid string, latitude float32, longitude float32) (Shop, error) {
	shop.UID = uid
	shop.Latitude = latitude
	shop.Longitude = longitude
	if err := db.DB.Create(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}

func (r ShopModel) Update(shop Shop, uid string, latitude float32, longitude float32) (Shop, error) {
	shop.UID = uid
	shop.Latitude = latitude
	shop.Longitude = longitude
	if err := db.DB.Updates(&shop).Error; err != nil {
		return shop, err
	}
	shop, err := r.GetByID(uid)
	if err != nil {
		return shop, err
	}
	return shop, nil
}

func (r ShopModel) Delete(shop Shop, uid string) (Shop, error) {
	if err := db.DB.Delete(&shop).Error; err != nil {
		return shop, err
	}
	return shop, nil
}
