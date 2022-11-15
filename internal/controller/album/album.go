// Package album implements methods associated
// to the application declared http routes.
package album

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"allsounds/pkg/repository/album"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type middleware interface {
	DB(handler gin.HandlerFunc) gin.HandlerFunc
}

// Middleware implement handler wrapper methods
type Middleware struct{}

func (m Middleware) DB(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", album.Album{})
		handler(context)
	}
}

func AddRoutes(router *gin.Engine, m middleware) *gin.Engine {
	router.GET("/album", Search)
	router.GET("/album/:id", m.DB(GetById))

	return router
}

// Search getAlbums responds with the list of all albums as JSON.
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

	var data []model.Album
	// If a query string has been passed, search albums by title, else fetch all
	if c.Query("query") != "" {
		albums := repository.Search(offset, limit, c.Query("query"), data, "title")
		c.IndentedJSON(http.StatusOK, albums)
	} else {
		albums := repository.FindAll(offset, limit, data, "title")
		c.IndentedJSON(http.StatusOK, albums)
	}
}

// GetById responds with a single album as JSON.
func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid id parameter")
		return
	}

	var data model.Album

	if db, ok := c.Get("db"); ok {
		err = db.(album.Repository).FindById(id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			c.Abort()
			log.Warn().Msgf("Bad request: can't fetch Album entity with id %v", id)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, data)
}
