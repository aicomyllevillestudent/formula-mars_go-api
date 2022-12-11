package models

import (
	"errors"
)

type Driver struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"size:255;not null" json:"name"`
	Races         []Race         `gorm:"many2many:race_drivers" json:"races"`
	Championships []Championship `gorm:"many2many:championship_drivers" json:"championships"`
}

func GetDrivers() ([]Driver, error) {
	var drivers []Driver

	if err := DB.Find(&drivers).Error; err != nil {
		return drivers, errors.New("drivers not found")
	}

	return drivers, nil
}

func (driver *Driver) AddDriver() (*Driver, error) {

	if err := DB.Create(&driver).Error; err != nil {
		return driver, errors.New("Race not found")
	}

	return driver, nil
}

func GetDriverById(id string) (Driver, error) {
	var driver Driver

	if err := DB.First(&driver, id).Error; err != nil {
		return driver, errors.New("Driver not found")
	}

	return driver, nil
}

func (driver *Driver) UpdateDriver(id string) (*Driver, error) {

	if err := DB.Model(&driver).Where(id).Updates(driver).Error; err != nil {
		return driver, err
	}

	return driver, nil
}

func (driver *Driver) DeleteDriver(id string) error {

	if err := DB.Where(id).Delete(&driver).Error; err != nil {
		return errors.New("Driver not deleted")
	}

	return nil
}
