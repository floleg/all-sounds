package track

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"allsounds/pkg/repository/track"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	var data []model.Track
	// If a query string has been passed, search artists by title, else fetch all
	if c.Query("query") != "" {
		tracks := repository.Search(offset, limit, c.Query("query"), data, "title")
		c.IndentedJSON(http.StatusOK, tracks)
	} else {
		tracks := repository.FindAll(offset, limit, data, "title")
		c.IndentedJSON(http.StatusOK, tracks)
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

	var data model.Track
	track, err := track.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msgf("request: can't fetch Track entity with id %v", id)

		return
	}

	c.IndentedJSON(http.StatusOK, track)
}
