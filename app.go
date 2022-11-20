package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type race struct {
	ID             int    `json:"id"`
	ChampionshipId int    `json:"championship_id"`
	Name           string `json:"name"`
	Date           string `json:"date"`
	Finished       bool   `json:"finished"`
}

var races = []race{
	{ID: 1, ChampionshipId: 1, Name: "The Martian Loop", Date: "2077-06-03 19:00", Finished: false},
	{ID: 2, ChampionshipId: 1, Name: "Interstellar Highway", Date: "2077-10-19 19:00", Finished: false},
}

func getRaces(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, races)
}

func getRaceByID(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not the right format"})
	}

	for _, a := range races {
		if a.ID == intId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "race not found"})
}

func addRace(c *gin.Context) {
	var newRace race

	if err := c.BindJSON(&newRace); err != nil {
		return
	}

	races = append(races, newRace)
	c.IndentedJSON(http.StatusCreated, newRace)
}

func main() {
	router := gin.Default()
	router.GET("/races", getRaces)
	router.GET("/races/:id", getRaceByID)
	router.POST("/races", addRace)

	if os.Getenv("PORT") != "" {
		router.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		router.Run("0.0.0.0:8080")
	}
}
