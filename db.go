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
	//DB.SetLogger(gorm.Logger{log.New(os.Stderr, "TRACE ", log.Ldate|log.Ltime|log.Lshortfile)})

	//f,_ := os.Create("log/"+time.Now().Format("2006-01-02")+"-sql.log")
	//DB.SetLogger(log.New(f, "GORM "+time.Now().Format("2006-01-02 15:04:05")+"\r\n", 0))

	// connection pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(200)
}

