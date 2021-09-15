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
	shop := new(controllers.ShopController)
	r.GET("/shop/:uid", shop.Get)
	r.POST("/management/shop", shop.Post)
	r.PATCH("/management/shop", shop.Patch)
	r.DELETE("management/shop", shop.Delete)
	search := new(controllers.SearchController)
	r.GET("/shop/search", search.Get)

	return r

}
