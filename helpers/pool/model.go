package pool

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var db *gorm.DB

func SetUp() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	fmt.Println(database)
	db, err := gorm.Open("mysql", database)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	defer db.Close()
	db.LogMode(true)
	log.Println("数据库连接成功")
}

func GetDB() *gorm.DB {
	return db
}
