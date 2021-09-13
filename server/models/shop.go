package models

type Shop struct {
	UID           string  `json:"uid" gorm:"primaryKey"`
	Name          string  `json:"name"`
	Image         string  `json:"image"`
	Address       string  `json:"address"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	SalesGoal     int     `json:"sales_goal"`
	Sales         int     `json:"sales"`
	Menus         []Menu
	RelatedGenres []RelatedGenre
}
