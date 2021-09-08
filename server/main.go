package main

import (
	"flag"

	"github.com/youngeek-0410/mottake/server/config"
	"github.com/youngeek-0410/mottake/server/db"
	"github.com/youngeek-0410/mottake/server/models"
	"github.com/youngeek-0410/mottake/server/router"
)

func main() {
	c := flag.String("config", "config", "config file")
	flag.Parse()
	config.Init(*c)
	db.Init()
	models.Init()
	r := router.NewRouter()
	r.Run(config.Config.Port)
}
