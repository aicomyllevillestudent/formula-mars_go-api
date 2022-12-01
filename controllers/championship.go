package controllers

import (
	"net/http"

	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/gin-gonic/gin"
)

type ChampionshipInput struct {
	Name string `json:"name"`
}

func GetChampionships(c *gin.Context) {
	championship, err := models.GetChampionships()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, championship)
}

func AddChampionship(c *gin.Context) {
	var input ChampionshipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	championship := models.Championship{}

	championship.Name = input.Name

	_, err := championship.AddChampionship()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, championship)
}

func GetChampionshipById(c *gin.Context) {
	id := c.Param("id")

	championship, err := models.GetChampionshipById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, championship)
}

func UpdateChampionship(c *gin.Context) {
	id := c.Param("id")

	var input ChampionshipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	championship := models.Championship{}

	championship.Name = input.Name

	_, err := championship.UpdateChampionship(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "name": championship.Name})

}

func DeleteChampionship(c *gin.Context) {
	id := c.Param("id")

	championship := models.Championship{}

	if err := championship.DeleteChampionship(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Championship has been deleted"})
}
