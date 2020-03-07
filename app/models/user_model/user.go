package user_model

import "gin-blog/helpers/pool"

type UserModel struct {

}

func GetUserALl()  {
	pool.GetDB().Select("").Attrs()
}