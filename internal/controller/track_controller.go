package controller

import (
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrackController struct{}

var trackRepository = new(repository.TrackRepository)

// responds with the list of all artists as JSON.
func (t TrackController) Search(c *gin.Context) {
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

	var data []model.Track
	// If a query string has been passed, search artists by title, else fetch all
	if c.Query("query") != "" {
		tracks := trackRepository.BaseRepo.Search(offset, limit, c.Query("query"), data, "title")
		c.IndentedJSON(http.StatusOK, tracks)
	} else {
		tracks := trackRepository.BaseRepo.FindAll(offset, limit, data, "title")
		c.IndentedJSON(http.StatusOK, tracks)
	}
}

// responds with a single artist as JSON.
func (t TrackController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	var data model.Track
	track, err := trackRepository.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, track)
}
