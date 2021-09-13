package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/db"
	"github.com/youngeek-0410/mottake/server/models"
)

type SearchController struct{}

const (
	numberOfShops  = 10                                 // 周辺の店をいくつまで表示するか
	searchRange    = 10                                 // 何km圏内の店を表示するか
	latitudeRange  = 0.0090133729745762 * searchRange   // kmを大体の緯度範囲に変換
	longitudeRange = 0.010966404715491394 * searchRange // kmを大体の経度範囲に変換
)

func (i SearchController) Get(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	var sortedShops []models.Shop
	// クエリの緯度経度から範囲内の店をソート
	sqlQuery := "SELECT * FROM shops WHERE ABS(@latitude - latitude) < @latrange AND ABS(@longitude - longitude) < @lonrange ORDER BY ABS(@latitude - latitude), ABS(@longitude - longitude) LIMIT @num"
	db.DB.Raw(sqlQuery, sql.Named("latitude", latitude), sql.Named("latrange", strconv.FormatFloat(latitudeRange, 'f', -1, 32)), sql.Named("longitude", longitude), sql.Named("lonrange", strconv.FormatFloat(longitudeRange, 'f', -1, 32)), sql.Named("num", numberOfShops)).Scan(&sortedShops)

	c.JSON(http.StatusOK, sortedShops)
}
