package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConn *gorm.DB
	err    error
)

func init() {
	dbConn, err = gorm.Open("mysql", "root:root@(35.194.206.151:3306)/stock?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbConn.LogMode(true)
}
