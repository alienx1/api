package service

import (
	"api/conf"
	"api/model/db"
)

func CreateBreeder(b *db.Breeder) (err error) {
	if err = conf.DB.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBreeder(b *db.Breeder, AnimalID string) (err error) {
	if err = conf.DB.Where("animal_id = ?", AnimalID).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}
