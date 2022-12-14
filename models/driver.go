package models

import (
	"errors"

	"gorm.io/gorm"
)

type Driver struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

type RaceDriver struct {
	RaceID   uint `gorm:"primaryKey" json:"race_id"`
	DriverID uint `gorm:"primaryKey" json:"driver_id"`
	Position int  `json:"position"`
	Laps     int  `json:"points"`
}

type DriverInRace struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	Laps     int    `json:"laps"`
}

func (driver *RaceDriver) BeforeSave(db *gorm.DB) error {
	driver.Position = 0
	driver.Laps = 0
	return nil
}

type ChampionshipDriver struct {
	ChampionshipID int `gorm:"primaryKey" json:"championship_id"`
	DriverID       int `gorm:"primaryKey" json:"driver_id"`
	Position       int `json:"position"`
	Points         int `json:"points"`
}

func GetDrivers() ([]Driver, error) {
	var drivers []Driver

	if err := DB.Find(&drivers).Error; err != nil {
		return drivers, errors.New("drivers not found")
	}

	return drivers, nil
}

func (driver *Driver) AddDriver() error {

	if err := DB.Create(&driver).Error; err != nil {
		return errors.New("Driver not added")
	}

	return nil
}

func GetDriverById(id string) (Driver, error) {
	var driver Driver

	if err := DB.First(&driver, id).Error; err != nil {
		return driver, errors.New("Driver not found")
	}

	return driver, nil
}

func (driver *Driver) UpdateDriver(id string) error {

	if err := DB.Model(&driver).Where(id).Updates(driver).Error; err != nil {
		return err
	}

	return nil
}

func (driver *Driver) DeleteDriver(id string) error {

	if err := DB.Where(id).Delete(&driver).Error; err != nil {
		return errors.New("Driver not deleted")
	}

	return nil
}
