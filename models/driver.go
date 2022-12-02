package models

import (
	"errors"
)

type Driver struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

func GetDrivers() ([]Driver, error) {
	var drivers []Driver

	if err := DB.Find(&drivers).Error; err != nil {
		return drivers, errors.New("drivers not found")
	}

	return drivers, nil
}
