package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 	model "one-piece-api/models"

// func getCharacter(c *gin.Context) {
// 	id := c.Param("id")
// 	var idInt int64
// 	_, err := fmt.Sscanf(id, "%d", &idInt)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		return
// 	}

// 	character, err := characterByID(idInt)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "character not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, character)
// }

func getDevilFruit(c *gin.Context) {
	id := c.Param("id")

	devilFruit, err := devilFruitByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "devil fruit not found"})
		return
	}

	c.JSON(http.StatusOK, devilFruit)
}

func router_setup() {
	router := gin.Default()

	router.GET("/devil_fruits/:id", getDevilFruit)
	// router.GET("/albums", getCharacters)

	router.Run("localhost:8080")
}
