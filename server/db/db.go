package db

import (
	"fmt"
	"log"

	"github.com/youngeek-0410/mottake/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	c := config.Config
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", c.DBConfig.Host, c.DBConfig.User, c.DBConfig.Name, c.DBConfig.Password)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
}
