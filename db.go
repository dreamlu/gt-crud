package deercoder

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	//database, initialize once
	DB, err = gorm.Open("mysql", GetConfigValue("db.user")+":"+GetConfigValue("db.password")+"@/"+GetConfigValue("db.name")+"?charset=utf8&parseTime=True&loc=Local")
	//defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Globally disable table names
	// use name replace names
	DB.SingularTable(true)
	// sql print console log
	DB.LogMode(true)
	//文件open
	nowTime := time.Now().Format("2006-01-02")
	filename := "log/"+nowTime+"-sql.log"
	var f *os.File
	_, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {//不存在，创建
			f,_ = os.Create(filename)
		}
	} else {
		f,_ = os.Open(filename)
	}
	DB.SetLogger(SetLogger(f))//打印错误信息和对应的sql

	// connection pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(200)
}

