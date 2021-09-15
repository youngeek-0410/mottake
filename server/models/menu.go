package models

import (
	"github.com/youngeek-0410/mottake/server/db"
	"gorm.io/plugin/soft_delete"
)

type Menu struct {
	ID        int                   `json:"id" gorm:"primaryKey"`
	ShopUID   string                `json:"shop_uid" gorm:"uniqueIndex:unique_menu"`
	Name      string                `json:"name" gorm:"uniqueIndex:unique_menu"`
	Price     int                   `json:"price"`
	Image     string                `json:"image"`
	Purchases []Purchase            `json:"purchases,omitempty"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:unique_menu"`
}

type MenuModel struct{}

func (m MenuModel) Create(shopUID string, menu Menu) (menuID int, err error) {
	menu.ShopUID = shopUID
	result := db.DB.Create(&menu)
	err = result.Error
	return menu.ID, err
}

func (m MenuModel) GetOneByID(shopUID string, menuID int) (menu Menu, err error) {
	result := db.DB.Where("id = ? AND shop_uid = ?", menuID, shopUID).First(&menu)
	if err = result.Error; err != nil {
		return menu, err
	}
	return menu, err
}

func (m MenuModel) All(shopUID string) (menus []Menu, err error) {
	result := db.DB.Where("shop_uid = ?", shopUID).Find(&menus)
	if err = result.Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (m MenuModel) Delete(shopUID string, menuID int) error {
	result := db.DB.Where("id = ? AND shop_uid = ?", menuID, shopUID).Delete(&Menu{})
	return result.Error
}

func (m MenuModel) Update(shopUID string, menuID int, menu Menu) error {
	result := db.DB.Model(&Menu{}).Where("id = ? AND shop_uid = ?", menuID, shopUID).Updates(&menu)
	return result.Error
}
