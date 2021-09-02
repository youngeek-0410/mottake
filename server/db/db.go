package db

import (
	"fmt"
	"log"

	"github.com/youngeek-0410/mottake/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	var err error
	c := config.Config
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", c.DbConfig.Host, c.DbConfig.User, c.DbConfig.Name, c.DbConfig.Password)
	Db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

}
