package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/gin-gonic/gin"
)

type RaceInput struct {
	ChampionshipID uint      `json:"championship_id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	Finished       bool      `json:"finished"`
}

func GetRaces(c *gin.Context) {
	r, err := models.GetRaces()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

func GetRaceByID(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not the right format"})
	}

	r, err := models.GetRaceByID(uint(intId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

func AddRace(c *gin.Context) {
	var input RaceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := models.Race{}

	r.ChampionshipID = input.ChampionshipID
	r.Name = input.Name
	r.Date = input.Date
	r.Finished = input.Finished

	err := r.AddRace()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Created")
}

func UpdateRace(c *gin.Context) {

	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not the right format"})
	}

	var input RaceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := models.Race{}

	r.ChampionshipID = input.ChampionshipID
	r.Name = input.Name
	r.Date = input.Date
	r.Finished = input.Finished

	error := r.UpdateRace(uint(intId))

	r.ID = uint(intId)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, "Updated")
}

func DeleteRace(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not the right format"})
	}

	r := models.Race{}

	if err := r.DeleteRace(uint(intId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}
