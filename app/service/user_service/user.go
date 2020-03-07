package user_service

import (
	"gin-blog/app/models"
)

type UserService struct {
	ID         int
	Account    string
	Password   string
	Status     uint
	CreateTime int
	UpdateTime int
}

func GetUserAll() []models.UserModel {
	result := models.GetUserALl()
	return result
}

func GetOne(id int) (user models.UserModel) {
	return models.GetOneById(id)
}