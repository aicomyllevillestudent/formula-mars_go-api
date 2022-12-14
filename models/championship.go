package models

import (
	"errors"
)

type Championship struct {
	ID      uint                   `gorm:"primaryKey" json:"id"`
	Name    string                 `gorm:"size:255;not null" json:"name"`
	Drivers []DriverInChampionship `gorm:"many2many:championship_drivers" json:"drivers"`
}

func GetChampionships() ([]Championship, error) {
	var championships []Championship

	if err := DB.Find(&championships).Error; err != nil {
		return championships, errors.New("Championship not found")
	}

	return championships, nil
}

func (championship *Championship) AddChampionship() error {

	if err := DB.Create(&championship).Error; err != nil {
		return errors.New("Race not found")
	}

	if err := AddDriversToChampionship(&championship.ID); err != nil {
		return err
	}

	return nil
}

func AddDriversToChampionship(id *uint) error {
	var drivers = []ChampionshipDriver{{ChampionshipID: *id, DriverID: 1}, {ChampionshipID: *id, DriverID: 2}, {ChampionshipID: *id, DriverID: 3}, {ChampionshipID: *id, DriverID: 4}, {ChampionshipID: *id, DriverID: 5}, {ChampionshipID: *id, DriverID: 6}}

	if err := DB.Create(&drivers).Error; err != nil {
		return err
	}

	return nil
}

func GetChampionshipById(id string) (Championship, error) {

	var championship Championship
	var drivers []Driver
	var championshipDrivers []ChampionshipDriver
	var result []DriverInChampionship

	if err := DB.First(&championship, id).Error; err != nil {
		return championship, errors.New("Championship not found")
	}

	if err := DB.Joins("JOIN championship_drivers ON championship_drivers.driver_id = drivers.id").Where("championship_drivers.championship_id = ?", id).Find(&drivers).Error; err != nil {
		return championship, err
	}

	DB.Where("championship_drivers.championship_id = ?", id).Find(&championshipDrivers)

	for _, driver := range drivers {
		result = append(result, DriverInChampionship{ID: driver.ID, Name: driver.Name})
	}

	for _, driver := range championshipDrivers {
		for i, res := range result {
			if res.ID == driver.DriverID {
				result[i].Points = driver.Points
			}
		}
	}

	championship.Drivers = result

	return championship, nil
}

func (championship *Championship) UpdateChampionship(id string) error {

	if err := DB.Model(&championship).Where(id).Updates(championship).Error; err != nil {
		return err
	}

	return nil
}

func (championship *Championship) DeleteChampionship(id string) error {

	if err := DB.Where(id).Delete(&championship).Error; err != nil {
		return errors.New("Championship not deleted")
	}

	return nil
}
