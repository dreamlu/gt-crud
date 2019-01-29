// author:  dreamlu
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

// dbcrud
var db = deercoder.DbCrud{
	Model: User{},		// model
	Table:"user",		// table name
}

// get user, by id
func (c *User)GetById(id string) interface{} {

	var user User	// not use *User
	db.ModelData = &user
	return db.GetByID(id)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User)GetBySearch(params map[string][]string) interface{} {
	var users []*User
	db.ModelData = &users
	return db.GetBySearch(params)
}

// delete user, by id
func (c *User)Delete(id string) interface{} {

	return db.Delete(id)
}

// update user
func (c *User)Update(params map[string][]string) interface{} {

	return db.Update(params)
}

// create user
func (c *User)Create(params map[string][]string) interface{} {

	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	return db.Create(params)
}
