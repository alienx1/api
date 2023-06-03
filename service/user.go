package service

import (
	"api/conf"
	"api/model"
	"api/model/db"
)

func CreateUser(u *db.User) (err error) {
	if err = conf.DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
func GetTableUser(u *[]model.User) (err error) {
	if err = conf.DB.Find(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetTableUserByUser(u *[]model.User, user string) (err error) {
	if err = conf.DB.Where("username like ?", user+"%").Error; err != nil {
		return err
	}
	return nil
}

func GetTableUserByName(u *[]model.User, name string) (err error) {
	if err = conf.DB.Where("name like ?", name+"%").Error; err != nil {
		return err
	}
	return nil
}

func GetTableUserByPhone(u *[]model.User, phone string) (err error) {
	if err = conf.DB.Where("phone like ?", phone+"%").Error; err != nil {
		return err
	}
	return nil
}

func GetTableUserByIDCard(u *[]model.User, id_card string) (err error) {
	if err = conf.DB.Where("id_card like ?", id_card+"%").Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(u *model.User, user string, id string) (err error) {
	if err = conf.DB.Where("username = ? AND id = ?", user, id).Table("users").Update(map[string]interface{}{
		"username":    u.Username,
		"name":        u.Name,
		"password":    u.Password,
		"birthday":    u.Birthday,
		"ethnic":      u.Ethnic,
		"nationality": u.Nationality,
		"religion":    u.Religion,
		"address":     u.Address,
		"phone":       u.Phone,
		"email":       u.Email,
		"duration":    u.Duration,
		"id":          u.ID,
		"credit":      u.Credit,
		"status":      u.Status,
		"update_time": u.UpdateTime,
		"update_by":   u.UpdateBy,
	}).Error; err != nil {
		return err
	}
	return nil
}
func UpdateStatus(u *model.User, user string) (err error) {
	if err = conf.DB.Where("username = ? ", user).Table("users").UpdateColumns(map[string]interface{}{
		"status":      u.Status,
		"update_time": u.UpdateTime,
		"update_by":   u.UpdateBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDuration(u *model.User, user string) (err error) {
	if err = conf.DB.Where("username = ?", user).Table("users").UpdateColumns(map[string]interface{}{
		"duration":    u.Duration,
		"update_time": u.UpdateTime,
		"update_by":   u.UpdateBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(u *model.User, user string) (err error) {
	if err = conf.DB.Where("username = ?", user).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
func FindSingleUser(u *model.User, user string) (err error) {
	if err = conf.DB.Where("username = ?", user).First(&u).Error; err != nil {
		return err
	}
	return nil
}
