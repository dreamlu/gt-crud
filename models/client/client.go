package client

import (
	"demo/models"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
)

// Client
type Client struct {
	models.AdminCom
	Name string `gorm:"type:varchar(30);" json:"name" valid:"required,len=2-20"` // 昵称
	// BirthDate time.CDate `json:"birth_date" gorm:"type:date"`
	// 请注意, 多账号的唯一性, 需要和AdminID一起建立唯一性
	// 如: 复制models.AdminCom中AdminID到这里 将"UNIQUE_INDEX:openid已存在"加入到AdminID中, 同时修改models.AdminCom为models.ModelCom
	//Openid     string     `json:"openid" gorm:"varchar(30);UNIQUE_INDEX:openid已存在"` // openID
	//HeadImg    string     `json:"head_img"` // 头像
}

var crud = gt.NewCrud()

// get data, by id
func (c *Client) Get(params cmap.CMap) (data Client, err error) {
	crud.Params(gt.Model(Client{}), gt.Data(&data))
	if err = crud.GetByData(params).Error(); err != nil {
		return
	}
	return
}

// get data, limit and search
func (c *Client) Search(params cmap.CMap) (datas []*Client, pager result.Pager, err error) {
	crud.Params(gt.Model(Client{}), gt.Data(&datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		return nil, pager, cd.Error()
	}
	return datas, cd.Pager(), nil
}

// delete data, by id
func (c *Client) Delete(id interface{}) error {

	return crud.Params(gt.Model(Client{})).Delete(id).Error()
}

// update data
func (c *Client) Update(data *Client) error {

	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	return nil
}

// create data
func (c *Client) Create(data *Client) (*Client, error) {

	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		return nil, err
	}
	return data, nil
}
