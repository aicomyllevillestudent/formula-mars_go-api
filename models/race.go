package models

import (
	"errors"
	"time"
)

type Race struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	ChampionshipId uint      `gorm:"not null" json:"championshipId"`
	Name           string    `gorm:"size:255;not null" json:"name"`
	Date           time.Time `gorm:"size:255;not null" json:"date"`
	Finished       bool      `gorm:"not null" json:"finished"`
}

func GetRaces() ([]Race, error) {
	var r []Race

	if err := DB.Find(&r).Error; err != nil {
		return r, errors.New("Race not found")
	}

	return r, nil
}

func GetRaceByID(uid uint) (Race, error) {

	var r Race

	if err := DB.First(&r, uid).Error; err != nil {
		return r, errors.New("Race not found")
	}

	return r, nil
}

func (r *Race) AddRace() (*Race, error) {
	if err := DB.Create(&r).Error; err != nil {
		return r, err
	}

	return r, nil
}

func (r *Race) UpdateRace(id uint) (*Race, error) {

	if err := DB.Model(&r).Where(id).Updates(r).Error; err != nil {
		return r, err
	}

	return r, nil
}

func (r *Race) DeleteRace(id uint) error {
	var dr Race

	res := DB.Where(id).Delete(&dr)

	if res.RowsAffected == 0 {
		return errors.New("Race not deleted")
	}

	return nil
}
