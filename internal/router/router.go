package router

import (
	"allsounds/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	album := new(controller.AlbumController)

	// Album entity routes declarations
	router.GET("/album", album.GetAlbums)
	router.GET("/album/:id", album.GetAlbumById)

	return router
}
