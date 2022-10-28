package controller

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistController struct{}

var artistRepository = new(repository.ArtistRepository)

// Search responds with the list of all artists as JSON.
func (a ArtistController) Search(c *gin.Context) {
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

	var data []model.Artist
	// If a query string has been passed, search artists by title, else fetch all
	if c.Query("query") != "" {
		artists := artistRepository.BaseRepo.Search(offset, limit, c.Query("query"), data, "name")
		c.IndentedJSON(http.StatusOK, artists)
	} else {
		artists := artistRepository.BaseRepo.FindAll(offset, limit, data, "name")
		c.IndentedJSON(http.StatusOK, artists)
	}
}

// GetById responds with a single artist as JSON.
func (a ArtistController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid id parameter")
		return
	}

	var data model.Artist
	artist, err := artistRepository.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msgf("Bad request: can't fetch Artist entity with id %v", id)
		return
	}

	c.IndentedJSON(http.StatusOK, artist)
}
