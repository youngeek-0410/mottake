package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

type MenuController struct{}

var menuModel = new(models.MenuModel)

func (m MenuController) Create(c *gin.Context) {
	var menu models.Menu
	uid := getUID(c)
	if err := c.ShouldBindJSON(&menu); err != nil {
		_ = c.Error(errInvalidJSONRequest).SetType(gin.ErrorTypePublic)
		return
	}

	menuID, err := menuModel.Create(uid, menu)
	if err != nil {
		_ = c.Error(errCouldNotCreateMenu).SetType(gin.ErrorTypePublic)
		return
	}
	menu, _ = menuModel.GetOneByID(uid, menuID)
	c.JSON(http.StatusOK, menu)
}

func (m MenuController) All(c *gin.Context) {
	shopUID := c.Param("uid")
	menus, err := menuModel.All(shopUID)
	if err != nil {
		_ = c.Error(errCouldNotGetMenus).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, menus)
}

func (m MenuController) One(c *gin.Context) {
	s := c.Param("menuID")
	shopUID := c.Param("uid")
	menuID, err := strconv.Atoi(s)
	if err != nil {
		_ = c.Error(errInvalidMenuID).SetType(gin.ErrorTypePublic)
		return
	}
	menu, err := menuModel.GetOneByID(shopUID, menuID)
	if err != nil {
		_ = c.Error(errCouldNotGetMenu).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, menu)
}

func (m MenuController) Delete(c *gin.Context) {
	uid := getUID(c)
	s := c.Param("menuID")
	menuID, err := strconv.Atoi(s)
	if err != nil {
		_ = c.Error(errInvalidMenuID).SetType(gin.ErrorTypePublic)
		return
	}
	err = menuModel.Delete(uid, menuID)
	if err != nil {
		_ = c.Error(errCouldNotDeleteMenu).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusNoContent, nil)

}

func (m MenuController) Update(c *gin.Context) {
	var menu models.Menu
	uid := getUID(c)
	s := c.Param("menuID")
	menuID, err := strconv.Atoi(s)
	if err != nil {
		_ = c.Error(errInvalidMenuID).SetType(gin.ErrorTypePublic)
		return
	}
	err = json.NewDecoder(c.Request.Body).Decode(&menu)
	if err != nil {
		_ = c.Error(errInvalidJSONRequest).SetType(gin.ErrorTypePublic)
		return
	}
	err = menuModel.Update(uid, menuID, menu)
	if err != nil {
		_ = c.Error(errCouldNotUpdateMenu).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
