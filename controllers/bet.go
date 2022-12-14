package controllers

import (
	"net/http"

	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/aicomylleville/formula-mars_go-api/utils/token"

	"github.com/gin-gonic/gin"
)

type BetInput struct {
	RaceID   int     `json:"race_id"`
	DriverId int     `json:"driver_id"`
	Amount   float64 `json:"amount"`
}

func GetBets(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bets, err := models.GetBets(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bets)
}

func AddBet(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input BetInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Wallet -= input.Amount

	u.UpdateWallet()

	bet := models.Bet{}

	bet.UserID = user_id
	bet.RaceID = input.RaceID
	bet.DriverId = input.DriverId
	bet.Amount = input.Amount

	error := bet.AddBet(u)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, "Created")
}
