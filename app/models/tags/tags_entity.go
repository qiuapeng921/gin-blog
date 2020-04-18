package tags

type Entity struct {
	Id          int    `orm:"id,primary"  json:"id"`          //
	Name        string `orm:"name"        json:"name"`        // 标签名称
	Description string `orm:"description" json:"description"` // 标签描述
	Status      int    `orm:"status"      json:"status"`      // 0 正常 1 删除
	CreateTime  int    `orm:"create_time" json:"create_time"` //
	UpdateTime  int    `orm:"update_time" json:"update_time"` //
}
