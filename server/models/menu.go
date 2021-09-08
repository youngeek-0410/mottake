package models

type Menu struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ShopUID   string `json:"shop_uid"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Image     string `json:"image"`
	Purchases []Purchase
}
