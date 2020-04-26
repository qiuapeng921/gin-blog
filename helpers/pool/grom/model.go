package grom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type Mysql struct {
	conn *gorm.DB
}

func SetUpOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	var mysql Mysql
	mysql.conn, err = gorm.Open("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	mysql.conn.SingularTable(true)
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	mysql.conn.DB().SetMaxIdleConns(maxIdle)
	mysql.conn.DB().SetMaxOpenConns(maxOpen)
	mysql.conn.LogMode(true)
}

func GetInstance() *Mysql {
	return &Mysql{}
}

func (m *Mysql) GetConn() *gorm.DB {
	return m.conn
}