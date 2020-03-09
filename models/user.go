// @author  dreamlu
package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/time"
)

// User
type User struct {
	ID         uint       `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
	Name       string     `gorm:"varchar(30)" json:"name" valid:"required,len=2-20"`         // 昵称
	Createtime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"createtime"` // 创建时间自动生成
	//Openid     string     `json:"openid" gorm:"varchar(30);UNIQUE_INDEX:openid已存在"` // openID
	//Headimg    string     `json:"headimg"` // 头像
}

var crud = gt.NewCrud(
	gt.Model(User{}),
	gt.Table("user"),
)

// get data, by id
func (c *User) GetByID(id string) interface{} {

	var data User // not use *User
	crud.Params(gt.Data(&data))
	if err := crud.GetByID(id).Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err.Error())
	}
	return result.GetSuccess(data)
}

// get data, limit and search
// clientPage 1, everyPage 10 default
func (c *User) GetBySearch(params map[string][]string) interface{} {
	var datas []*User
	crud.Params(gt.Data(&datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		//log.Log.Error(err.Error())
		return result.GetError(cd.Error())
	}
	return result.GetSuccessPager(datas, cd.Pager())
}

// delete data, by id
func (c *User) Delete(id string) interface{} {

	if err := crud.Delete(id).Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update data
func (c *User) Update(data *User) interface{} {

	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create data
func (c *User) Create(data *User) interface{} {

	// create time
	//(*data).Createtime = time2.CTime(time.Now())

	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}

// update data
func (c *User) UpdateForm(params map[string][]string) interface{} {

	if err := crud.UpdateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create data
func (c *User) CreateForm(params map[string][]string) interface{} {

	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))

	if err := crud.CreateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
