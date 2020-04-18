package users

type Entity struct {
	Id         uint   `orm:"id,primary"      json:"id"`          //
	Username   string `orm:"username,unique" json:"username"`    // 用户名
	Password   string `orm:"password"        json:"password"`    // 密码
	Phone      string `orm:"phone,unique"    json:"phone"`       // 手机号
	Email      string `orm:"email"           json:"email"`       // 邮箱
	Avatar     string `orm:"avatar"          json:"avatar"`      // 头像
	Status     uint   `orm:"status"          json:"status"`      // 状态 0:删除 1:正常 2:禁用
	CreateTime uint   `orm:"create_time"     json:"create_time"` // 创建时间
	UpdateTime uint   `orm:"update_time"     json:"update_time"` // 更新时间
}

func (Entity) TableName() string {
	return "users"
}