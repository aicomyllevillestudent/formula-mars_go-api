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
	championships, err := models.GetChampionships()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, championships)
}

func AddChampionship(c *gin.Context) {
	var input ChampionshipInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	championship := models.Championship{}

	championship.Name = input.Name

	err := championship.AddChampionship()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Created")
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

	err := championship.UpdateChampionship(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Updated")

}

func DeleteChampionship(c *gin.Context) {
	id := c.Param("id")

	championship := models.Championship{}

	if err := championship.DeleteChampionship(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}
