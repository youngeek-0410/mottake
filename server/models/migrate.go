package models

import "github.com/youngeek-0410/mottake/server/db"

func Init() {
	db.Db.AutoMigrate(&Customer{}, &Shop{}, &Genre{}, &FavoriteGenre{}, &RelatedGenre{}, &Receipt{}, &Purchase{}, &Menu{})
}
