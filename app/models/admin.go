package models

type AdminModel struct {
	Id         uint   `orm:"id,primary"  json:"id"`          // 用户id
	Username   string `orm:"username"    json:"username"`    // 用户名
	Password   string `orm:"password"    json:"password"`    // 密码
	Phone      string `orm:"phone"       json:"phone"`       // 手机号
	Status     int    `orm:"status"      json:"status"`      // 状态:0正常 1禁用
	LoginIp    string `orm:"login_ip"    json:"login_ip"`    // 最后登录ip
	LoginTime  uint   `orm:"login_time"  json:"login_time"`  // 最后登录时间
	CreateTime uint   `orm:"create_time" json:"create_time"` // 添加时间
	UpdateTime uint   `orm:"update_time" json:"update_time"` // 修改时间
}

func (AdminModel) TableName() string {
	return "admins"
}
