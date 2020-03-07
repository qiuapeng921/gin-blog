package models

import (
	"gin-blog/helpers/pool/grom"
)

type UserModel struct {
	ID         int    `json:"id"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Status     uint   `json:"status"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

func (UserModel) TableName() string {
	return "user"
}

func GetUserALl() ([]*UserModel, error) {
	db := grom.GetConn()
	var users []*UserModel
	err := db.Select("id,account,status").Find(&users).Error
	return users, err
}

func GetOneById(id int) (*UserModel, error) {
	db := grom.GetConn()
	var users *UserModel
	err := db.Where("account=?", "admin1").Find(&users).Error
	return users, err
}
