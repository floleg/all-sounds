package controller

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
	"allsounds/pkg/repository"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userRepository = new(repository.UserRepository)

// Search responds with the list of all artists as JSON.
func (u UserController) Search(c *gin.Context) {
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

// GetById responds with a single artist as JSON.
func (u UserController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid id parameter")
		return
	}

	var data model.User
	user, err := userRepository.FindById(id, data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msgf("Bad request: can't fetch User entity with id %s", id)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// AppendUserTrack responds with a single user as JSON.
func (u UserController) AppendUserTrack(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid userId parameter")
		return
	}

	trackId, err := strconv.Atoi(c.Param("trackId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		c.Abort()
		log.Warn().Msg("Bad request: invalid trackId parameter")
		return
	}

	user := model.User{}
	userRepository.BaseRepo.FindById(userId, &user)

	track := model.Track{}
	trackRepository.BaseRepo.FindById(trackId, &track)

	user.Tracks = append(user.Tracks, track)

	db.DBCon.Omit("Track").Save(&user)

	c.IndentedJSON(http.StatusOK, user)
}
