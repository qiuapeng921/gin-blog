package grom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var conn *gorm.DB

func SetUpOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	db, err := gorm.Open("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.LogMode(true)
	conn = db
	fmt.Println("mysql连接成功")
}

func GetConn() *gorm.DB {
	return conn
}