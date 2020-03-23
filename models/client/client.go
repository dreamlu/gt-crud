package client

import (
	"demo/models"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
)

// Client
type Client struct {
	models.ModelCom
	Name string `gorm:"type:varchar(30)" json:"name" valid:"required,len=2-20"` // 昵称
	//Openid     string     `json:"openid" gorm:"varchar(30);UNIQUE_INDEX:openid已存在"` // openID
	//Headimg    string     `json:"headimg"` // 头像
}

var crud = gt.NewCrud(
	gt.Model(Client{}),
)

// get data, by id
func (c *Client) GetByID(id string) interface{} {

	var data Client // not use *Client
	crud.Params(gt.Data(&data))
	if err := crud.GetByID(id).Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.CError(err)
	}
	return result.GetSuccess(data)
}

// get data, limit and search
// clientPage 1, everyPage 10 default
func (c *Client) GetBySearch(params map[string][]string) interface{} {
	var datas []*Client
	crud.Params(gt.Data(&datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		//log.Log.Error(err.Error())
		return result.CError(cd.Error())
	}
	return result.GetSuccessPager(datas, cd.Pager())
}

// delete data, by id
func (c *Client) Delete(id string) interface{} {

	if err := crud.Delete(id).Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeDelete, result.MsgDelete)
}

// update data
func (c *Client) Update(data *Client) interface{} {

	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.CError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create data
func (c *Client) Create(data *Client) interface{} {

	// create time
	//(*data).Createtime = time2.CTime(time.Now())

	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		//log.Log.Error(err.Error())
		return result.CError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}

// update data
func (c *Client) UpdateForm(params map[string][]string) interface{} {

	if err := crud.UpdateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeUpdate, result.MsgUpdate)
}

// create data
func (c *Client) CreateForm(params map[string][]string) interface{} {

	//params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))

	if err := crud.CreateForm(params); err != nil {
		//log.Log.Error(err.Error())
		return result.GetError(err)
	}
	return result.GetMapData(result.CodeCreate, result.MsgCreate)
}
