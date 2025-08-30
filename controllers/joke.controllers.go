package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"main/services"
	"math/rand"
	"net/http"
	"strconv"
)

func GetJokes(c *gin.Context) {
	jokes, err := services.GetJokes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jokes)
	return
}

func CreateJoke(c *gin.Context) {
	var newJoke models.Joke
	err := c.ShouldBindJSON(&newJoke)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := services.CreateJoke(newJoke)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
	return
}

func UpdateJoke(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid joke id"})
		return
	}
	var updatedJoke models.Joke
	if err := c.ShouldBindJSON(&updatedJoke); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := services.UpdateJoke(uint(id), updatedJoke)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
	return
}

func DeleteJoke(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid joke id"})
		return
	}
	
	err = services.DeleteJoke(uint(id))
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusNoContent, gin.H{"message": "Joke deleted successfully"})
}

func GetRandomJoke(c *gin.Context) {
	jokes, err := services.GetJokes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(jokes) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No jokes found"})
		return
	}
	randomIndex := rand.Intn(len(jokes))
	c.JSON(http.StatusOK, jokes[randomIndex])
	return
}
