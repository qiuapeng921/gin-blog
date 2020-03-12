package models

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
