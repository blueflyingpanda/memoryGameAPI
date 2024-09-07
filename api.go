package main

import (
	_ "InterfacesAPI/docs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// ListUsers godoc
// @Summary List all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func ListUsers(c *gin.Context) {
	users := GetAllUsers()
	c.IndentedJSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get a user by username
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} User
// @Failure 404 {object} map[string]string
// @Router /users/{username} [get]
func GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := GetUserByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// ListPlayers godoc
// @Summary List all players
// @Tags players
// @Accept  json
// @Produce  json
// @Success 200 {array} Player
// @Router /players [get]
func ListPlayers(c *gin.Context) {
	players := GetAllPlayers()
	c.IndentedJSON(http.StatusOK, players)
}

// GetPlayer godoc
// @Summary Get a player by login
// @Tags players
// @Accept  json
// @Produce  json
// @Param login path string true "Login"
// @Success 200 {object} Player
// @Failure 404 {object} map[string]string
// @Router /players/{login} [get]
func GetPlayer(c *gin.Context) {
	login := c.Param("login")

	player, err := GetPlayerByLogin(login)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "player not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, player)
}

type PlayerRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AddPlayer godoc
// @Summary Add a new player
// @Tags players
// @Accept json
// @Produce json
// @Param player body PlayerRequest true "Player data"
// @Success 200 {object} Player
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /players [post]
func AddPlayer(c *gin.Context) {
	var json PlayerRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !IsValidSHA256Hash(json.Password) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "You gotta be kidding. Did you really just sent an unhashed password? 😂 Try SHA256"})
		return
	}

	player, err := CreatePlayer(json.Login, json.Password)
	if err != nil {
		var statusCode int

		if errors.As(err, &ErrPlayerExists) || errors.As(err, &ErrNoSuchUser) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}

		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"login": player.Login,
		"score": player.Score,
	})
}

type ScoreRequest struct {
	Score uint `json:"score" binding:"required"`
}

// UpdatePlayer godoc
// @Summary Update a player's score
// @Tags players
// @Accept json
// @Produce json
// @Param login path string true "Login"
// @Param body body ScoreRequest true "Score data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /players/{login} [put]
func UpdatePlayer(c *gin.Context) {
	login := c.Param("login")

	var json ScoreRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	score, err := SetPlayerScore(login, json.Score)
	if err != nil {
		var statusCode int

		if errors.As(err, &ErrNoSuchPlayer) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"login": login,
		"score": score,
	})
}

// Ping godoc
// @Summary Ping test endpoint
// @Tags ping
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"ping": "pong"})
}

func initAPI(port int) {
	router := gin.Default()

	router.GET("/ping", Ping)

	router.GET("/users", ListUsers)
	router.GET("/users/:username", GetUser)

	router.GET("/players", ListPlayers)
	router.GET("/players/:login", GetPlayer)
	router.POST("/players", AddPlayer)
	router.PUT("/players/:login", UpdatePlayer)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(fmt.Sprintf(":%d", port))
}
