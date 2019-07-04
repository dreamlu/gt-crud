// @author  dreamlu
package models

import (
	"demo/util/log"
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
	if err := db_json.GetByID(id); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	return result.GetSuccess(user)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearchJ(params map[string][]string) interface{} {
	var users []*User
	db_json.ModelData = &users

	pager, err := db.GetBySearch(params)
	if err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetSuccessPager(users, pager)
}

// delete user, by id
func (c *User) DeleteJ(id string) interface{} {

	if err := db_json.Delete(id); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update user
func (c *User) UpdateJ(data *User) interface{} {

	if err := db_json.Update(data); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create user
func (c *User) CreateJ(data *User) interface{} {

	// create time
	(*data).Createtime = time2.CTime(time.Now())

	if err := db_json.Create(data); err != nil {
		log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
