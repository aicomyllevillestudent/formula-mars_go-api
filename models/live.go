package models

import (
	"errors"
	"fmt"
)

type Live struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	RaceID int    `json:"-"`
	Race   Race   `gorm:"foreignKey:RaceID;references:ID" json:"race"`
	Link   string `json:"link"`
}

func GetLive() (Live, error) {
	var live Live
	var drivers []Driver
	var raceDrivers []RaceDriver
	var result []DriverInRace

	if err := DB.Preload("Race").Last(&live).Error; err != nil {
		return live, errors.New("no live race found")
	}

	if err := DB.Joins("JOIN race_drivers ON race_drivers.driver_id = drivers.id").Where("race_drivers.race_id = ?", live.Race.ID).Find(&drivers).Error; err != nil {
		return live, err
	}

	DB.Where("race_drivers.race_id = ?", live.Race.ID).Find(&raceDrivers)

	for _, driver := range drivers {
		result = append(result, DriverInRace{ID: driver.ID, Name: driver.Name})
	}

	for _, driver := range raceDrivers {
		for i, res := range result {
			if res.ID == driver.DriverID {
				result[i].Position = driver.Position
				result[i].Laps = driver.Laps
			}
		}
	}

	live.Race.Drivers = result

	return live, nil
}

func (live *Live) AddLive() error {

	if err := DB.Create(&live).Error; err != nil {
		return errors.New("live race not added")
	}

	return nil
}

func (driver *DriverInRace) UpdateLive(id uint) error {

	live, _ := GetLive()
	race := live.Race
	var raceDriver = RaceDriver{}

	raceDriver.Laps = driver.Laps
	raceDriver.Position = driver.Position

	if err := DB.Model(&RaceDriver{}).Where("race_id = ? and driver_id = ?", race.ID, id).Updates(&raceDriver).Error; err != nil {
		return err
	}

	return nil
}

func (live *Live) DeleteLive() error {

	var race Race
	var drivers []RaceDriver
	var winner uint
	var bets []Bet

	liveResult, _ := GetLive()

	race.Finished = true

	if err := race.UpdateRace(liveResult.Race.ID); err != nil {
		return err
	}

	DB.Where("race_drivers.race_id = ?", liveResult.Race.ID).Find(&drivers)

	for _, driver := range drivers {
		points := GetPoints(driver.Position)
		if driver.Position == 1 {
			winner = driver.DriverID
		}
		DB.Model(&RaceDriver{}).Where("race_drivers.race_id = ? and race_drivers.driver_id = ?", liveResult.Race.ID, driver.DriverID).UpdateColumn("points", points)
		DB.Model(&ChampionshipDriver{}).Where("championship_drivers.championship_id = ? and championship_drivers.driver_id = ?", liveResult.Race.ChampionshipID, driver.DriverID).UpdateColumn("points", points)
	}

	DB.Where("race_id = ? and driver_id = ?", live.Race.ID, winner).Find(&bets)

	for _, bet := range bets {
		var user User
		DB.Where("id = ?", bet.UserID).Find(&user)
		user.Wallet += bet.Amount * 2
		fmt.Println(user.Wallet)
		fmt.Println(bet.Amount * 2)
		fmt.Println(user.Wallet + bet.Amount*2)
		DB.Model(&User{}).Where("id = ?", bet.UserID).UpdateColumn("wallet", user.Wallet)
	}

	if err := DB.Last(&live).Delete(&live).Error; err != nil {
		return errors.New("live race not deleted")
	}

	return nil
}

func GetPoints(position int) int {
	var points = 0

	switch position {
	case 1:
		points = 15
	case 2:
		points = 10
	case 3:
		points = 5
	case 4:
		points = 2
	default:
		points = 0
	}

	return points
}
