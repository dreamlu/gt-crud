package applet

import (
	"demo/models"
	"fmt"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
)

// 小程序相关 模型
type Applet struct {
	models.AdminCom
	Appid     string `json:"appid" gorm:"type:varchar(50);UNIQUE_INDEX:appid已存在"` // appid
	Secret    string `json:"secret" gorm:"type:varchar(50)"`                      // secret
	MchID     string `json:"mch_id" gorm:"type:varchar(50)"`                      // 商户号
	PaySecret string `json:"pay_secret" gorm:"type:varchar(50)"`                  // 商户api秘钥
	// 证书
	AppCert string `json:"app_cert"` // apiApplet_cert.pem
	AppKey  string `json:"app_key"`  // apiApplet_key.pem
	// logo
	//Logo string `json:"logo"` // logo
}

var crud = gt.NewCrud(
	gt.Model(Applet{}),
)

// 多账号使用
// get data, by id
func (c *Applet) GetByAdminID(admin_id uint64) error {

	//var data Applet // not use *Applet
	sql := fmt.Sprintf("select %s from applet where admin_id = ?", gt.GetColSQL(Applet{}))
	cd := crud.Params(gt.Data(&c)).Select(sql, admin_id).Single()
	if err := cd.Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	//c = &data
	return nil
}

// get data, by id
func (c *Applet) GetByAppid(appid string) error {

	//var data Applet // not use *Applet
	sql := fmt.Sprintf("select %s from applet where appid = ?", gt.GetColSQL(Applet{}))
	cd := crud.Params(gt.Data(&c)).Select(sql, appid).Single()
	if err := cd.Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	//c = &data
	return nil
}

// get data, by id
func (c *Applet) Get(params cmap.CMap) (data Applet, err error) {
	crud.Params(gt.Data(&data))
	if err = crud.GetByData(params).Error(); err != nil {
		return
	}
	return
}

// get data, limit and search
// AppletPage 1, everyPage 10 default
func (c *Applet) Search(params cmap.CMap) (datas []*Applet, pager result.Pager, err error) {
	//var datas []*Applet
	crud.Params(gt.Data(&datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		//log.Log.Error(err.Error())
		return nil, pager, cd.Error()
	}
	return datas, cd.Pager(), nil
}

// delete data, by id
func (c *Applet) Delete(id string) error {

	return crud.Delete(id).Error()
}

// update data
func (c *Applet) Update(data *Applet) (*Applet, error) {

	crud.Params(gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return nil, err
	}
	return data, nil
}

// create data
func (c *Applet) Create(data *Applet) (*Applet, error) {

	crud.Params(gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		return nil, err
	}
	return data, nil
}
