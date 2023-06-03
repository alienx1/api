package service

import (
	"api/conf"
	"api/model/db"
)

func CreateAnimalImg(ai *db.AnimalImage) (err error) {
	if err = conf.DB.Create(&ai).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnimalImg(ai *db.AnimalImage, AnimalID string) (err error) {
	if err = conf.DB.Where("animal_id = ?", AnimalID).Delete(&ai).Error; err != nil {
		return err
	}
	return nil
}
