package service

import (
	"api/conf"
	"api/model"
	"api/model/db"
)

func CreateAnimal(a *db.Animal) (err error) {
	if err := conf.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnimal(a *db.Animal, AnimalID string) (err error) {
	if err = conf.DB.Where("animal_id = ?", AnimalID).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimal(res *[]model.Animal) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	ORDER BY a.animal_id  
`).Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimalByAnimalID(res *[]model.Animal, animal_id string) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	WHERE a.animal_id like ?
	ORDER BY a.animal_id  
`, animal_id+"%").Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimalOwnerName(res *[]model.Animal, owner_name string) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	WHERE a.owner_name like ?
	ORDER BY a.animal_id  
`, owner_name+"%").Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimalName(res *[]model.Animal, name string) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	WHERE a.name like ?
	ORDER BY a.animal_id  
`, name+"%").Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimalType(res *[]model.Animal, Type string) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	WHERE a.type like ?
	ORDER BY a.animal_id  
`, Type+"%").Scan(&res).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAnimalByAnimalIDWithPet(res *model.Animal, animal_id string) (err error) {
	if err = conf.DB.Raw(`
	SELECT 	a.animal_id , a.birthday ,a.ear_number ,a.name ,a.owner_name ,a.type ,a.gender, 
		a.color ,a.defect ,b.breeder1  ,b.breeder2 ,ai.image1 ,ai.image2 ,ai.image3 ,ai.image4 ,ai.image5,
		a.create_time ,a.create_by ,a.update_time ,a.update_by
	FROM  animals a 
	Join breeders b ON	b.animal_id =a.animal_id 
	JOIN animal_images ai ON a.animal_id = ai.animal_id 
	WHERE a.animal_id = ?
`, animal_id).Scan(&res).Error; err != nil {
		return err
	}
	return nil
}
