package controller

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct{}

var r = new(repository.Repository)

// getAlbums responds with the list of all albums as JSON.
func (a AlbumController) Search(c *gin.Context) {
	if c.Query("offset") == "" || c.Query("limit") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	var data []model.Album
	// If a query string has been passed, search albums by title, else fetch all
	if c.Query("query") != "" {
		albums := r.Search(offset, limit, c.Query("query"), data)
		c.IndentedJSON(http.StatusOK, albums)
	} else {
		albums := r.FindAll(offset, limit, data)
		c.IndentedJSON(http.StatusOK, albums)
	}
}

// getAlbums responds with a single album as JSON.
func (a AlbumController) GetAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	var data model.Album
	album := r.FindById(id, &data)

	c.IndentedJSON(http.StatusOK, album)
}
