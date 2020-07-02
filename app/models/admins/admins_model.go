package admins

import "gin-blog/helpers/pool/grom"

func GetOne(id int64) (admins Entity) {
	grom.GetOrm().Where("id = ?", id).First(&admins)
	return
}
