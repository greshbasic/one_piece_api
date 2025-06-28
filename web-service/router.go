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

func getAllDevilFruitsHandler(c *gin.Context) {
	dfs, err := getAllDevilFruits()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dfs)
}

func getLocationHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location ID"})
		return
	}

	location, err := getLocationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

func getAllLocationsHandler(c *gin.Context) {
	locations, err := getAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func getArtifactHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid artifact ID"})
		return
	}

	artifact, err := getArtifactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, artifact)
}

func getAllArtifactsHandler(c *gin.Context) {
	artifacts, err := getAllArtifacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, artifacts)
}

func router_setup() {
	router := gin.Default()

	router.GET("/characters", getAllCharactersHandler)
	router.GET("/characters/:id", getCharacterHandler)

	router.GET("/devil_fruits", getAllDevilFruitsHandler)
	router.GET("/devil_fruits/:id", getDevilFruitHandler)

	router.GET("/locations", getAllLocationsHandler)
	router.GET("/locations/:id", getLocationHandler)

	router.GET("/artifacts", getAllArtifactsHandler)
	router.GET("/artifacts/:id", getArtifactHandler)

	router.Run("localhost:8080")
}
