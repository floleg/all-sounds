// Package router instantiates a gin engine and
// declares REST API routes, associating controller
// layer business methods.
package router

import (
	"allsounds/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Album entity routes declarations
	album := new(controller.AlbumController)
	router.GET("/album", album.Search)
	router.GET("/album/:id", album.GetById)

	// Artist entity routes declarations
	artist := new(controller.ArtistController)
	router.GET("/artist", artist.Search)
	router.GET("/artist/:id", artist.GetById)

	// Track entity routes declarations
	track := new(controller.TrackController)
	router.GET("/track", track.Search)
	router.GET("/track/:id", track.GetById)

	// Track entity routes declarations
	user := new(controller.UserController)
	router.GET("/user", user.Search)
	router.GET("/user/:id", user.GetById)
	router.POST("/user/:userId/track/:trackId", user.AppendUserTrack)

	return router
}
