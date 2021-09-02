package models

type Menu struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	ShopUid   string `json:"shop_uid"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	Purchases []Purchase
}
