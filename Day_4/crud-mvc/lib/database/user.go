package database

import (
	"crud-mvc/config"
	middleware "crud-mvc/middelware"
	"crud-mvc/models"
	"fmt"
)

func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	fmt.Println("Masuk", user.Email, user.Password, user.ID)
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUser(userId int) (interface{}, error) {
	var user models.User
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

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
