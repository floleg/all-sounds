// Package router instantiates a gin engine and
// declares REST API routes, associating controller
// layer business methods.
package router

import (
	"allsounds/internal/controller/album"
	"allsounds/internal/controller/artist"
	"allsounds/internal/controller/track"
	"allsounds/internal/controller/user"
	"github.com/gin-gonic/gin"
)

// NewRouter instantiates and return a gin Engine router with its declared routes and controller methods
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Album entity routes declarations
	album.AddRoutes(router, album.Middleware{})

	// Artist entity routes declarations
	artist.AddRoutes(router)

	// Track entity routes declarations
	track.AddRoutes(router)

	// Track entity routes declarations
	user.AddRoutes(router)

	return router
}
