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

	return router
}
