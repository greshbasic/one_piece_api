package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCharacterHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	character, err := getCharacterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}

func getAllCharactersHandler(c *gin.Context) {
	characters, err := getAllCharacters()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func getDevilFruitHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid devil fruit ID"})
		return
	}

	df, err := getDevilFruitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, df)
}

func router_setup() {
	router := gin.Default()

	router.GET("/characters", getAllCharactersHandler)
	router.GET("/characters/:id", getCharacterHandler)

	router.GET("/devil_fruits/:id", getDevilFruitHandler)

	router.Run("localhost:8080")
}
