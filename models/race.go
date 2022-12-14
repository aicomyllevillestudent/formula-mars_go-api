package models

import (
	"errors"
	"time"
)

type Race struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ChampionshipId uint           `gorm:"not null" json:"championshipId"`
	Name           string         `gorm:"size:255;not null" json:"name"`
	Date           time.Time      `gorm:"size:255;not null" json:"date"`
	Finished       bool           `gorm:"not null" json:"finished"`
	Drivers        []DriverInRace `gorm:"many2many:race_drivers;" json:"drivers"`
}

func GetRaces() ([]Race, error) {
	var r []Race

	if err := DB.Find(&r).Error; err != nil {
		return r, errors.New("Race not found")
	}

	return r, nil
}

func GetRaceByID(uid uint) (Race, error) {

	var race Race
	var drivers []Driver
	var raceDrivers []RaceDriver
	var result []DriverInRace

	if err := DB.First(&race, uid).Error; err != nil {
		return race, errors.New("Race not found")
	}

	if err := DB.Joins("JOIN race_drivers ON race_drivers.driver_id = drivers.id").Where("race_drivers.race_id = ?", uid).Find(&drivers).Error; err != nil {
		return race, err
	}

	DB.Where("race_drivers.race_id = ?", uid).Find(&raceDrivers)

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

	race.Drivers = result

	return race, nil
}

func (r *Race) AddRace() error {

	if err := DB.Create(&r).Error; err != nil {
		return err
	}

	if err := AddDriversToRace(&r.ID); err != nil {
		return err
	}

	return nil
}

func AddDriversToRace(id *uint) error {
	var drivers = []RaceDriver{{RaceID: *id, DriverID: 1}, {RaceID: *id, DriverID: 2}, {RaceID: *id, DriverID: 3}, {RaceID: *id, DriverID: 4}, {RaceID: *id, DriverID: 5}, {RaceID: *id, DriverID: 6}}

	if err := DB.Create(&drivers).Error; err != nil {
		return err
	}

	return nil
}

func (r *Race) UpdateRace(id uint) error {

	if err := DB.Model(&r).Where(id).Updates(r).Error; err != nil {
		return err
	}

	return nil
}

func (r *Race) DeleteRace(id uint) error {
	var dr Race

	res := DB.Where(id).Delete(&dr)

	if res.RowsAffected == 0 {
		return errors.New("Race not deleted")
	}

	return nil
}
