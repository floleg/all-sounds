// Package controller implements methods associated
// to the application declared http routes.
package controller

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct{}

var albumRepository = new(repository.AlbumRepository)

// Search getAlbums responds with the list of all albums as JSON.
func (a AlbumController) Search(c *gin.Context) {
	if c.Query("offset") == "" || c.Query("limit") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: missing offset or limit parameter")
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid offset parameter")
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid limit parameter")
		return
	}

	var data []model.Album
	// If a query string has been passed, search albums by title, else fetch all
	if c.Query("query") != "" {
		albums := albumRepository.BaseRepo.Search(offset, limit, c.Query("query"), data, "title")
		c.IndentedJSON(http.StatusOK, albums)
	} else {
		albums := albumRepository.BaseRepo.FindAll(offset, limit, data, "title")
		c.IndentedJSON(http.StatusOK, albums)
	}
}

// GetById responds with a single album as JSON.
func (a AlbumController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid id parameter")
		return
	}

	var data model.Album

	album, err := albumRepository.FindById(id, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msgf("Bad request: can't fetch Album entity with id %v", id)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
