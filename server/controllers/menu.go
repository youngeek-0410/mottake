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
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}

	menuID, err := menuModel.Create(uid, menu)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, errCouldNotCreateMenu})
		return
	}
	menu, _ = menuModel.GetOneByID(uid, menuID)
	c.JSON(http.StatusOK, menu)
}

func (m MenuController) All(c *gin.Context) {
	shopUID := c.Param("uid")
	menus, err := menuModel.All(shopUID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errCouldNotGetMenus})
		return
	}
	c.JSON(http.StatusOK, menus)
}

func (m MenuController) One(c *gin.Context) {
	s := c.Param("menuID")
	shopUID := c.Param("uid")
	menuID, err := strconv.Atoi(s)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidMenuID})
		return
	}
	menu, err := menuModel.GetOneByID(shopUID, menuID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errCouldNotGetMenu})
		return
	}
	c.JSON(http.StatusOK, menu)
}

func (m MenuController) Delete(c *gin.Context) {
	uid := getUID(c)
	s := c.Param("menuID")
	menuID, err := strconv.Atoi(s)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errInvalidMenuID})
		return
	}
	err = menuModel.Delete(uid, menuID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, errCouldNotDeleteMenu})
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
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errInvalidMenuID})
		return
	}
	err = json.NewDecoder(c.Request.Body).Decode(&menu)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}
	err = menuModel.Update(uid, menuID, menu)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, errCouldNotUpdateMenu})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
