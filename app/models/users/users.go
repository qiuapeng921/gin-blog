package users

import (
	"fmt"
)

func GetUserByUsername(username string) (Entity, error) {
	return FindOne(fmt.Sprintf("username = %s", username))
}

func GetUserById(id uint) (Entity, error) {
	return FindOne(fmt.Sprintf("id = %d", id))
}
