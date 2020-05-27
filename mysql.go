package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func connectDB(user, password, dbPath, database string) (*gorm.DB, error) {
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbPath, database)
	return gorm.Open("mysql", sqlUrl)
}

func ConnectMysql(dbPath, password string) (*gorm.DB, error) {
	return connectDB("gongwei", password, dbPath, "gongwei")
}
