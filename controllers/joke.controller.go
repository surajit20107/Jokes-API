package controller

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"main/services"
	"net/http"
	"strconv"
	"math/rand"
)

func GetJokes(c *gin.Context) {
	jokes := services.GetJokes()
	c.JSON(http.StatusOK, jokes)
	return
}

func GetJokesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid joke id"})
	}
	joke, err := services.GetJokesById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, joke)
}

func CreateJoke(c *gin.Context) {
	var newJoke models.Joke
	if err := c.ShouldBindJSON(&newJoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}
	created := services.CreateJoke(newJoke.Content, newJoke.Author)
	c.JSON(http.StatusCreated, created)
	return
}

func UpdateJoke(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid joke id"})
	}
	var updatedJoke models.Joke
	if err := c.ShouldBindJSON(&updatedJoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := services.UpdateJoke(id, updatedJoke)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, updated)
	return
}

func DeleteJoke(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid joke id"})
	}
	err = services.DeleteJoke(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Joke deleted successfully"})
	return
}

func GetRandomJoke(c *gin.Context) {
	jokes := services.GetJokes()
	randomIndex := rand.Intn(len(jokes))
	c.JSON(http.StatusOK, jokes[randomIndex])
}
