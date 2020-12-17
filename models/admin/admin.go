package admin

import (
	"demo/models"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/util"
)

// 多账号管理 模型
type Admin struct {
	models.ModelCom
	Name     string `gorm:"type:varchar(30);uniqueIndex:账号已存在" json:"name"` // 名称
	Password string `json:"password" gorm:"type:varchar(100)"`              // 密码
	Role     *int8  `json:"role" gorm:"type:tinyint(2);DEFAULT:0"`          // 0默认(超级管理员),1管理员
}

var crud = gt.NewCrud(
	gt.Model(Admin{}),
)

// get data, by id
func (c *Admin) Get(params cmap.CMap) (data Admin, err error) {
	crud.Params(gt.Data(&data))
	if err = crud.GetByData(params).Error(); err != nil {
		return
	}
	return
}

// get data, limit and search
// AdminPage 1, everyPage 10 default
func (c *Admin) Search(params cmap.CMap) (datas []*Admin, pager result.Pager, err error) {
	//var datas []*Admin
	crud.Params(gt.Data(&datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		//log.Log.Error(err.Error())
		return nil, pager, cd.Error()
	}
	return datas, cd.Pager(), nil
}

// delete data, by id
func (c *Admin) Delete(id string) error {

	return crud.Delete(id).Error()
}

// update data
func (c *Admin) Update(data *Admin) error {

	data.Password = util.AesEn(data.Password)
	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	return nil
}

// create data
func (c *Admin) Create(data *Admin) (*Admin, error) {

	data.Password = util.AesEn(data.Password)
	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		return nil, err
	}
	return data, nil
}
