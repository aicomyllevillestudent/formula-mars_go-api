package controllers

import (
	"net/http"

	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/gin-gonic/gin"
)

type DriverInput struct {
	Name string `json:"name"`
}

func GetDrivers(c *gin.Context) {
	drivers, err := models.GetDrivers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drivers)
}
