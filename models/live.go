package models

import "errors"

type Live struct {
	ID      uint     `gorm:"primaryKey" json:"id"`
	RaceID  int      `json:"race_id"`
	Drivers []Driver `json:"drivers"`
	Link    string   `json:"link"`
}

func GetLive() (Live, error) {
	var live Live

	if err := DB.Find(&live).Last(&live).Error; err != nil {
		return live, errors.New("no live race found")
	}

	return live, nil
}

func (live *Live) AddLive() error {

	if err := DB.Create(&live).Error; err != nil {
		return errors.New("live race not added")
	}

	return nil
}

func (live *Live) UpdateLive() error {

	if err := DB.Find(&live).Last(&live).Updates(&live).Error; err != nil {
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
