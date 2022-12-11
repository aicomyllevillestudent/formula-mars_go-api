package models

import (
	"errors"
)

type Championship struct {
	ID      uint     `gorm:"primaryKey" json:"id"`
	Name    string   `gorm:"size:255;not null" json:"name"`
	Drivers []Driver `gorm:"many2many:championship_drivers" json:"drivers"`
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

	return nil
}

func GetChampionshipById(id string) (Championship, error) {
	var championship Championship

	if err := DB.First(&championship, id).Preload("Drivers").Error; err != nil {
		return championship, errors.New("Championship not found")
	}

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
