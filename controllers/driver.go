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

func AddDriver(c *gin.Context) {
	var input DriverInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver := models.Driver{}

	driver.Name = input.Name

	err := driver.AddDriver()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Created")
}

func GetDriverById(c *gin.Context) {
	id := c.Param("id")

	driver, err := models.GetDriverById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}

func UpdateDriver(c *gin.Context) {
	id := c.Param("id")

	var input DriverInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver := models.Driver{}

	driver.Name = input.Name

	err := driver.UpdateDriver(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Updated")

}

func DeleteDriver(c *gin.Context) {
	id := c.Param("id")

	driver := models.Driver{}

	if err := driver.DeleteDriver(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}
