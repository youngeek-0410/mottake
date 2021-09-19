package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/models"
)

type FavoriteGenreController struct{}

var favoriteGenreModel = new(models.FavoriteGenreModel)

func (i FavoriteGenreController) GetByID(c *gin.Context) {
	uid := getUID(c)
	genre, err := favoriteGenreModel.GetByID(uid)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errNotFound})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func (i FavoriteGenreController) Save(c *gin.Context) {
	var genre models.FavoriteGenre
	if err := c.BindJSON(&genre); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}

	uid := getUID(c)
	genre.CustomerUID = uid
	returnedGenre, err := favoriteGenreModel.Save(genre)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, errCouldNotUpdateGenre})
		return
	}
	c.JSON(http.StatusOK, returnedGenre)
}

type RelatedGenreController struct{}

var relatedGenreModel = new(models.RelatedGenreModel)

func (i RelatedGenreController) GetByID(c *gin.Context) {
	uid := getUID(c)
	genre, err := relatedGenreModel.GetByID(uid)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errNotFound})
		return
	}
	c.JSON(http.StatusOK, genre)
}

func (i RelatedGenreController) Save(c *gin.Context) {
	var genre models.RelatedGenre
	if err := c.BindJSON(&genre); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, errInvalidJSONRequest})
		return
	}

	uid := getUID(c)
	genre.ShopUID = uid
	returnedGenre, err := relatedGenreModel.Save(genre)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusInternalServerError, errCouldNotUpdateGenre})
		return
	}
	c.JSON(http.StatusOK, returnedGenre)
}

type GenreController struct{}

var GenreModel = new(models.GenreModel)

func (i GenreController) GetAll(c *gin.Context) {
	genre, err := GenreModel.GetAll()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusNotFound, errNotFound})
		return
	}
	c.JSON(http.StatusOK, genre)
}
