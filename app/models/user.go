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
func GetOneById(id int) (users UserModel) {
	db := grom.GetConn()
	db.Where("account=?", "admin1").Find(&users)
	return
}

func GetUserALl() (users []UserModel) {
	db := grom.GetConn()
	db.Select("id,account,status").Find(&users)
	return users
}
