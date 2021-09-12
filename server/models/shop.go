package models

import "github.com/youngeek-0410/mottake/server/geocoder"

type Shop struct {
	UID     string `json:"uid" gorm:"primaryKey"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Address string `json:"address"`
	// TODO: 構造体をどうにかせねば
	Coordinate    geocoder.Coordinate `json:"coordinate"`
	SalesGoal     int                 `json:"sales_goal"`
	Sales         int                 `json:"sales"`
	Menus         []Menu
	RelatedGenres []RelatedGenre
}
