package deercoder

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	//database, initialize once
	DB, err = gorm.Open("mysql", GetConfigValue("db.user")+":"+GetConfigValue("db.password")+"@"+GetConfigValue("db.password")+"/"+GetConfigValue("db.name")+"?charset=utf8&parseTime=True&loc=Local")
	//defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Globally disable table names
	// use name replace names
	DB.SingularTable(true)
	// sql print console log
	// or print sql err to file
	LogMode("debug") // or sqlErr

	// connection pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(200)
}

