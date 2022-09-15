package database

import (
	"crud-mvc/config"
	"crud-mvc/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e

	}
	return users, nil
}

func CreateUser(user models.User) (interface{}, error) {

	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e

	}
	return user, nil
}

func UpdateUser(user models.User) (interface{}, error) {

	if e := config.DB.Where("id = ?", user.ID).Updates(&user).Error; e != nil {
		return nil, e

	}
	return user, nil
}

func DeleteUser(user models.User) (interface{}, error) {

	if e := config.DB.Where("id = ?", user.ID).Delete(&user).Error; e != nil {
		return nil, e

	}
	return user, nil
}
