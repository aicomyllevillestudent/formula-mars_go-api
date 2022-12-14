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
	ID       uint `json:"id"`
	Position int  `json:"position"`
	Laps     int  `json:"laps"`
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
	live.Link = "https://stream.aronbuffel.be/hls/race.m3u8"

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

	driver := models.DriverInRace{}

	driver.Position = input.Position
	driver.Laps = input.Laps

	err := driver.UpdateLive(input.ID)

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
