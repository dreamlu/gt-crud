package models

import (
	"github.com/dreamlu/deercoder-gin"
	"time"
)

/*user model*/
type User struct {
	ID         uint               `json:"id" gorm:"primary_key"`
	Name       string             `json:"name"`
	Createtime deercoder.JsonTime `json:"createtime"` //maybe you like util.JsonDate
}

// get user, by id
func (c *User)GetById(id string) interface{} {

	deercoder.DB.AutoMigrate(&User{})
	var user = User{}
	return deercoder.GetDataById(&user, id)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User)GetBySearch(args map[string][]string) interface{} {
	//相当于注册类型,https://github.com/jinzhu/gorm/issues/857
	//db.DB.AutoMigrate(&User{})
	//var users = []*User{}
	var users []*User
	return deercoder.GetDataBySearch(User{}, &users, "user", args) //匿名User{}
}

// delete user, by id
func (c *User)Delete(id string) interface{} {

	return deercoder.DeleteDataByName("user", "id", id)
}

// update user
func (c *User)Update(args map[string][]string) interface{} {

	return deercoder.UpdateData("user", args)
}

// create user
func (c *User)Create(args map[string][]string) interface{} {

	args["createtime"] = append(args["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	return deercoder.CreateData("user", args)
}
