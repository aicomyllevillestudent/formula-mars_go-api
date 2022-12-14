package models

import "errors"

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

	if err := DB.Find(&live).Last(&live).Delete(&live).Error; err != nil {
		return errors.New("live race not deleted")
	}

	return nil
}
