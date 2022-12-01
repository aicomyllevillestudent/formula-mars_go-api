package models

import (
	"errors"
)

type Championship struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

func GetChampionships() ([]Championship, error) {
	var championship []Championship

	if err := DB.Find(&championship).Error; err != nil {
		return championship, errors.New("Championship not found")
	}

	return championship, nil
}

func (championship *Championship) AddChampionship() (*Championship, error) {

	if err := DB.Create(&championship).Error; err != nil {
		return championship, errors.New("Race not found")
	}

	return championship, nil
}

func GetChampionshipById(id string) (Championship, error) {
	var championship Championship

	if err := DB.First(&championship, id).Error; err != nil {
		return championship, errors.New("Championship not found")
	}

	return championship, nil
}

func (championship *Championship) UpdateChampionship(id string) (*Championship, error) {

	if err := DB.Model(&championship).Where(id).Updates(championship).Error; err != nil {
		return championship, err
	}

	return championship, nil
}

func (championship *Championship) DeleteChampionship(id string) error {

	if err := DB.Where(id).Delete(&championship).Error; err != nil {
		return errors.New("Championship not deleted")
	}

	return nil
}
