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

	shop := new(controllers.ShopController)
	r.GET("/shop/:uid", shop.GetByID)
	r.POST("/management/shop", shop.Create)
	r.PATCH("/management/shop", shop.Update)
	r.DELETE("management/shop", shop.Delete)
	search := new(controllers.SearchController)
	r.GET("/shop/search", search.Get)

	receipt := new(controllers.ReceiptController)
	r.GET("/user/receipt", receipt.All)
	r.POST("/management/shop/receipt", receipt.Create)

	favoriteGenre := new(controllers.FavoriteGenreController)
	r.GET("/user/genre", favoriteGenre.GetByID)
	r.PUT("/user/genre", favoriteGenre.Save)

	relatedGenre := new(controllers.RelatedGenreController)
	r.GET("/shop/:uid/genre", relatedGenre.GetByID)
	r.PUT("/management/shop/genre", relatedGenre.Save)

	genre := new(controllers.GenreController)
	r.GET("/genre", genre.GetAll)

	customer := new(controllers.CustomerController)
	r.GET("/user", customer.Show)
	r.POST("/user", customer.Create)
	r.PATCH("/user", customer.Update)
	r.DELETE("/user", customer.Destroy)
	return r
}
