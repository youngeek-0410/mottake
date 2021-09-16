package models

import (
	"github.com/youngeek-0410/mottake/server/db"
	"gorm.io/gorm"
)

type Customer struct {
	Name          string        `json:"name" binding:"required"`
	UID           string        `json:"uid" gorm:"primaryKey"`
	FavoriteGenre FavoriteGenre `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CustomerModel struct{}

func (cus CustomerModel) GetOneByID(uid string) (Customer, error) {
	var customer Customer
	var result *gorm.DB

	result = db.DB.Where("uid = ?", uid).Preload("FavoriteGenre.Genre").First(&customer)

	if err := result.Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (cus CustomerModel) Create(uid string, customer Customer) (reCus Customer, err error) {
	customer.UID = uid
	err = db.DB.Create(&customer).Error
	reCus = customer
	return reCus, err
}

func (cus CustomerModel) Update(uid string, customer Customer) (reCus Customer, err error) {
	err = db.DB.Where("uid = ?", uid).Updates(&customer).Error
	reCus = customer
	return reCus, err
}

func (cus CustomerModel) Destroy(uid string) error {
	var customer Customer
	err := db.DB.Where("uid = ?", uid).Delete(&customer).Error
	return err
}
