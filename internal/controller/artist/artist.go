package artist

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"allsounds/pkg/repository/artist"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/artist", Search)
	router.GET("/artist/:id", GetById)

	return router
}

// Search responds with the list of all artists as JSON.
func Search(c *gin.Context) {
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
		artists := repository.Search(offset, limit, c.Query("query"), data, "name")
		c.IndentedJSON(http.StatusOK, artists)
	} else {
		artists := repository.FindAll(offset, limit, data, "name")
		c.IndentedJSON(http.StatusOK, artists)
	}
}

// GetById responds with a single artist as JSON.
func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid id parameter")
		return
	}

	var data model.Artist
	artistEntity, err := artist.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msgf("Bad request: can't fetch Artist entity with id %v", id)
		return
	}

	c.IndentedJSON(http.StatusOK, artistEntity)
}
