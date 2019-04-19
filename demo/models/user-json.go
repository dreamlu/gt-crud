// author:  dreamlu
package models

import (
	"github.com/dreamlu/deercoder-gin"
	"time"
)

///*user model*/
//type User struct {
//	ID         uint               `json:"id" gorm:"primary_key"`
//	Name       string             `json:"name"`
//	Createtime deercoder.JsonTime `json:"createtime"` //maybe you like util.JsonDate
//}

// dbcrud
var db_json = deercoder.DbCrudJ{
	Model: User{},		// model
	Table:"user",		// table name
}

// get user, by id
func (c *User)GetByIDJ(id string) interface{} {

	var user User	// not use *User
	db_json.ModelData = &user
	return db.GetByID(id)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User)GetBySearchJ(params map[string][]string) interface{} {
	var users []*User
	db_json.ModelData = &users
	return db.GetBySearch(params)
}

// delete user, by id
func (c *User)DeleteJ(id string) interface{} {

	return db_json.Delete(id)
}

// update user
func (c *User)UpdateJ(data *User) interface{} {

	return db_json.Update(data)
}

// create user
func (c *User)CreateJ(data *User) interface{} {

	// create time
	(*data).Createtime = deercoder.JsonTime(time.Now())

	return db_json.Create(data)
}
