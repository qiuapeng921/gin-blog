package admins

import "gin-blog/helpers/pool/grom"

func GetList() (admins []*Entity, count int64) {
	orm := grom.GetConn()
	orm.Limit(10).Offset(1).Find(&admins)
	count = orm.RowsAffected
	return admins, count
}
