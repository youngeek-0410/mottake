package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/config"
	"github.com/youngeek-0410/mottake/server/controllers"
	"github.com/youngeek-0410/mottake/server/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())
	if !(config.Config.Mode == config.ModeDevelop) {
		r.Use(middleware.Auth())
	} else {
		r.Use(middleware.Dummy())
	}

	index := new(controllers.IndexController)
	r.GET("/", index.Get)
	menu := new(controllers.MenuController)
	r.GET("/shop/:uid/menu", menu.All)
	r.GET("/shop/:uid/menu/:menuID", menu.One)
	r.POST("/management/shop/menu", menu.Create)
	r.PATCH("/management/shop/menu/:menuID", menu.Update)
	r.DELETE("/management/shop/menu/:menuID", menu.Delete)

	return r

}
