package db

import (
	"fmt"
	"deercoder-gin/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

var (
	DB *gorm.DB
)

func init() {
	var err error
	//数据库,全局初始化一次
	DB, err = gorm.Open("mysql", conf.GetConfigValue("db.user")+":"+conf.GetConfigValue("db.password")+"@/"+conf.GetConfigValue("db.name")+"?charset=utf8&parseTime=True&loc=Local")
	//defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	//全局禁用表名复数
	DB.SingularTable(true)
	//sql打印
	DB.LogMode(true)
	//连接池
	//最大闲置连接与打开连接
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(100)
}
