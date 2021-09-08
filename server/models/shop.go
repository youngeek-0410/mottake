package models

type Shop struct {
	Uid           string `json:"uid" gorm:"primaryKey"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Image         string `json:"image"`
	Address       string `json:"address"`
	SalesGoal     int    `json:"sales_goal"`
	Sales         int    `json:"sales"`
	Menus         []Menu
	RelatedGenres []RelatedGenre
}
