package controller

import (
	"allsounds/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct{}

var albumModel = new(model.Album)

// getAlbums responds with the list of all albums as JSON.
func (a AlbumController) GetAlbums(c *gin.Context) {
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

	albums := albumModel.FindAll(offset, limit)

	c.IndentedJSON(http.StatusOK, albums)
}