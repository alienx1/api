package service

import (
	"api/conf"
	"api/model"
	"api/model/db"
)

func GetTableAdmin(a *[]db.Admin) (err error) {
	if err = conf.DB.Find(&a).Error; err != nil {
		return err
	}
	return nil
}

func GetTableAdminByUser(a *[]db.Admin, user string) (err error) {
	if err = conf.DB.Where("username like ?", user+"%").Find(&a).Error; err != nil {
		return err
	}
	return nil
}
func GetTableAdminByName(a *[]db.Admin, name string) (err error) {
	if err = conf.DB.Where("name like ?", name+"%").Find(&a).Error; err != nil {
		return err
	}
	return nil
}

func CraeteAdmin(a *db.Admin) (err error) {
	if err = conf.DB.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdmin(a *model.Admin, user string) (err error) {
	if err = conf.DB.Table("admins").Where("username = ?", user).Updates(map[string]interface{}{
		"username":    a.Username,
		"name":        a.Name,
		"password":    a.Password,
		"rank":        a.Rank,
		"status":      a.Status,
		"create_time": a.CreateTime,
		"create_by":   a.CreateBy,
		"update_time": a.UpdateTime,
		"update_by":   a.UpdateBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(a *db.Admin, user string) (err error) {
	if err = conf.DB.Where("username = ?", user).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func FindSingleAdmin(a *model.Admin, user string) (err error) {
	if err = conf.DB.Where("username = ?", user).First(&a).Error; err != nil {
		return err
	}
	return nil
}
