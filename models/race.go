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
	if err := DB.DB().Ping(); err != nil {
		ConnectDatabase()
	}

	var r []Race

	if err := DB.Find(&r).Error; err != nil {
		return r, errors.New("Race not found")
	}

	if err := DB.Close(); err != nil {
		return r, errors.New("not closed")
	}

	return r, nil
}

func GetRaceByID(uid uint) (Race, error) {
	if err := DB.DB().Ping(); err != nil {
		ConnectDatabase()
	}

	var r Race

	if err := DB.First(&r, uid).Error; err != nil {
		return r, errors.New("Race not found")
	}

	if err := DB.Close(); err != nil {
		return r, errors.New("not closed")
	}

	return r, nil
}

func (r *Race) AddRace() (*Race, error) {
	if err := DB.DB().Ping(); err != nil {
		ConnectDatabase()
	}

	if err := DB.Create(&r).Error; err != nil {
		return r, err
	}

	if err := DB.Close(); err != nil {
		return r, errors.New("not closed")
	}

	return r, nil
}

func (r *Race) UpdateRace(id uint) (*Race, error) {
	if err := DB.DB().Ping(); err != nil {
		ConnectDatabase()
	}

	if err := DB.Model(&r).Where(id).Updates(r).Error; err != nil {
		return r, err
	}

	if err := DB.Close(); err != nil {
		return r, errors.New("not closed")
	}

	return r, nil
}

func (r *Race) DeleteRace(id uint) error {
	if err := DB.DB().Ping(); err != nil {
		ConnectDatabase()
	}

	var dr Race

	res := DB.Where(id).Delete(&dr)

	if res.RowsAffected == 0 {
		return errors.New("Race not deleted")
	}

	if err := DB.Close(); err != nil {
		return errors.New("not closed")
	}

	return nil
}
