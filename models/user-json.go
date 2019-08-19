// @author  dreamlu
package models

import (
	. "demo/util/global"
	der "github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/tool/result"
	time2 "github.com/dreamlu/go-tool/tool/type/time"
	"log"
	"time"
)

var crud = der.DBCrud{
	DBTool: DBTool,
	Param: &der.CrudParam{
		Model: User{}, // model
		Table: "user", // table name
	},
}

// get user, by id
func (c *User) GetByIDJ(id string) interface{} {

	var user User // not use *User
	crud.Param.ModelData = &user
	if err := crud.GetByID(id); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	log.Print("测试", crud.Param)
	return result.GetSuccess(user)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearchJ(params map[string][]string) interface{} {
	var users []*User
	crud.Param.ModelData = &users

	pager, err := crud.GetBySearch(params)
	if err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetSuccessPager(users, pager)
}

// delete user, by id
func (c *User) DeleteJ(id string) interface{} {

	if err := crud.Delete(id); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update user
func (c *User) UpdateJ(data *User) interface{} {

	if err := crud.Update(data); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create user
func (c *User) CreateJ(data *User) interface{} {

	// create time
	(*data).Createtime = time2.CTime(time.Now())

	if err := crud.Create(data); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
