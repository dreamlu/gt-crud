// @author  dreamlu
package models

import (
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/tool/result"
	time2 "github.com/dreamlu/go-tool/tool/type/time"
	"time"
)

///*user model*/
//type User struct {
//	ID         uint               `json:"id" gorm:"primary_key"`
//	Name       string             `json:"name"`
//	Createtime der.JsonTime `json:"createtime"` //maybe you like tool.JsonDate
//}

// dbcrud json
var db_json = der.DbCrudJ{
	Model: User{}, // model
	Table: "user", // table name
}

// get user, by id
func (c *User) GetByIDJ(id string) interface{} {

	var user User // not use *User
	db_json.ModelData = &user
	return db.GetByID(id)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearchJ(params map[string][]string) interface{} {
	var users []*User
	db_json.ModelData = &users

	pager, err := db.GetBySearch(params)
	if err != nil {
		return result.GetError(err)
	}
	return result.GetSuccessPager(users, pager)
}

// delete user, by id
func (c *User) DeleteJ(id string) interface{} {

	return db_json.Delete(id)
}

// update user
func (c *User) UpdateJ(data *User) interface{} {

	return db_json.Update(data)
}

// create user
func (c *User) CreateJ(data *User) interface{} {

	// create time
	(*data).Createtime = time2.CTime(time.Now())

	return db_json.Create(data)
}
