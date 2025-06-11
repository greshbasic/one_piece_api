package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDevilFruit(c *gin.Context) {
	id := c.Param("id")

	devilFruit, err := devilFruitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "devil fruit not found"})
		return
	}

	c.JSON(http.StatusOK, devilFruit)
}

func getDevilFruits(c *gin.Context) {
	devilFruits, err := allDevilFruits()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "devil fruits not found"})
		return
	}

	c.JSON(http.StatusOK, devilFruits)
}

func getCharacter(c *gin.Context) {
	id := c.Param("id")

	character, err := characterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "character not found"})
		return
	}

	c.JSON(http.StatusOK, character)
}

func getCharacters(c *gin.Context) {
	characters, err := allCharacters()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "devil fruits not found"})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func router_setup() {
	router := gin.Default()

	router.GET("/devil_fruits", getDevilFruits)
	router.GET("/devil_fruits/:id", getDevilFruit)

	router.GET("/characters", getCharacters)
	router.GET("/characters/:id", getCharacter)

	router.Run("localhost:8080")
}
