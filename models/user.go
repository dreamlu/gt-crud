// @author  dreamlu
package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	. "github.com/dreamlu/gt/tool/type/time"
	"time"
)

/*user model*/
type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name" valid:"required,len=2-20"`
	Createtime CTime  `json:"createtime"` //maybe you like util.JsonDate
}

var crud2 = gt.NewCrud(
	gt.Model(User{}),
	gt.Table("user"),
)

// get user, by id
func (c *User) GetByID(id string) interface{} {

	var user User // not use *User
	crud.Params(gt.Data(&user))
	if err := crud2.GetByID(id); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	return result.GetSuccess(user)
}

// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearch(params map[string][]string) interface{} {
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
func (c *User) Delete(id string) interface{} {

	if err := crud2.Delete(id); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update user
func (c *User) Update(params map[string][]string) interface{} {

	if err := crud2.UpdateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create user
func (c *User) Create(params map[string][]string) interface{} {

	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))

	if err := crud2.CreateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
