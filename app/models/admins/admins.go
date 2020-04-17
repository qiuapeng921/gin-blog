package admins

import "gin-blog/helpers/pool/grom"

func GetList(params ...map[string]interface{}) (admins []*Entity, count int64) {
	var model = grom.GetConn()
	if params != nil {
		for key, val := range params {
			model.Where(key, val)
		}
	}
	result := model.Limit(10).Offset(1).Find(&admins)
	count = result.RowsAffected
	return
}

func GetOne(id int64) (admins *Entity) {
	grom.GetConn().Where("id = ?", id).First(&admins)
	return
}