package controller

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userRepository = new(repository.UserRepository)

// responds with the list of all artists as JSON.
func (u UserController) Search(c *gin.Context) {
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

	var data []model.User
	// If a query string has been passed, search artists by title, else fetch all
	if c.Query("query") != "" {
		users := userRepository.BaseRepo.Search(offset, limit, c.Query("query"), data, "login")
		c.IndentedJSON(http.StatusOK, users)
	} else {
		users := userRepository.BaseRepo.FindAll(offset, limit, data, "login")
		c.IndentedJSON(http.StatusOK, users)
	}
}

// responds with a single artist as JSON.
func (u UserController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	var data model.User
	user, err := userRepository.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// responds with a single artist as JSON.
func (u UserController) AppendUserTrack(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	trackId, err := strconv.Atoi(c.Param("trackId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	user := model.User{}
	userRepository.BaseRepo.FindById(userId, &user)

	track := model.Track{}
	trackRepository.BaseRepo.FindById(trackId, &track)

	user.Tracks = append(user.Tracks, track)

	db.DBCon.Omit("Track").Save(&user)

	// err = userRepository.AppendUserTrack(&user, &track)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
