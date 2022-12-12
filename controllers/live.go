package controllers

import (
	"net/http"

	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/gin-gonic/gin"
)

type AddLiveInput struct {
	RaceID int    `json:"race_id"`
	Link   string `json:"link"`
}

type UpdateLiveInput struct {
	Drivers []models.Driver `json:"drivers"`
}

func GetLive(c *gin.Context) {

	live, err := models.GetLive()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, live)
}

func AddLive(c *gin.Context) {

	var input AddLiveInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	live := models.Live{}

	live.RaceID = input.RaceID
	live.Link = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

	err := live.AddLive()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Created")
}

func UpdateLive(c *gin.Context) {

	var input UpdateLiveInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	live := models.Live{}

	// live.Drivers = input.Drivers

	err := live.UpdateLive()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Updated")
}

func DeleteLive(c *gin.Context) {
	live := models.Live{}

	if err := live.DeleteLive(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}
