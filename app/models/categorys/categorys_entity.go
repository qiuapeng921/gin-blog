package categorys

type Entity struct {
	Id          int    `orm:"id,primary"  json:"id"`          //
	Name        string `orm:"name"        json:"name"`        // 分类名称
	Description string `orm:"description" json:"description"` // 分类描述
	Status      int    `orm:"status"      json:"status"`      // 0 正常 1 删除
	CreateTime  int    `orm:"create_time" json:"create_time"` // 创建时间
	UpdateTime  int    `orm:"update_time" json:"update_time"` // 修改时间
}

func (Entity) TableName() string {
	return "categorys"
}
