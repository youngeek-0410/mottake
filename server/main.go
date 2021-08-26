package main

import (
	"flag"
	"fmt"

	"github.com/youngeek-0410/mottake/server/config"
	"github.com/youngeek-0410/mottake/server/db"
	"github.com/youngeek-0410/mottake/server/router"
)

func main() {
	fmt.Println("Hello, Wd")
	c := flag.String("config", "config", "config file")
	flag.Parse()
	config.Init(*c)
	db.Init()
	r := router.NewRouter()
	r.Run(config.Config.Port)
}
