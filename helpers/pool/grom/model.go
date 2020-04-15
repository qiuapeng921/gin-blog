package grom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
)

var db *gorm.DB

func SetUpGOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	db, err = gorm.Open("mysql", database)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	db.SingularTable(true)
	maxIdle,_:= strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen,_:= strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.LogMode(true)
}

func GetConn() *gorm.DB {
	return db
}


var engine *xorm.Engine

func SetUpXOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	engine, err = xorm.NewEngine("mysql", database)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	maxIdle,_:= strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen,_:= strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxOpen)
	engine.ShowSQL(true)
}

func GetXOrmConn() *xorm.Engine {
	return engine
}
