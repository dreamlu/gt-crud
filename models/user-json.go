// @author  dreamlu
package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	time2 "github.com/dreamlu/gt/tool/type/time"
	"time"
)

var crud = gt.NewCrud(
	gt.Model(User{}),
	gt.Table("user"),
)

// get user, by id
func (c *User) GetByIDJ(id string) interface{} {

	var user User // not use *User
	crud.Params(gt.Data(&user))
	if err := crud.GetByID(id); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	gt.Logger().Print("测试", user)
	return result.GetSuccess(user)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearchJ(params map[string][]string) interface{} {
	var users []*User
	crud.Params(gt.Data(&users))

	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		//log.Log.Error(err.Error())
		return result.GetError(cd.Error())
	}
	return result.GetSuccessPager(users, cd.Pager())
}

// delete user, by id
func (c *User) DeleteJ(id string) interface{} {

	if err := crud.Delete(id).Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update user
func (c *User) UpdateJ(data *User) interface{} {

	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create user
func (c *User) CreateJ(data *User) interface{} {

	// create time
	(*data).Createtime = time2.CTime(time.Now())

	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
