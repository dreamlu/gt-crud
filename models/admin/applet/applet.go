package applet

import (
	"demo/models"
	"fmt"
	"github.com/dreamlu/gt"
)

// 小程序相关 模型
type Applet struct {
	models.AdminCom
	Appid     string `json:"appid" gorm:"type:varchar(50);uniqueIndex:appid已存在"` // appid
	Secret    string `json:"secret" gorm:"type:varchar(50)"`                     // secret
	MchID     string `json:"mch_id" gorm:"type:varchar(50)"`                     // 商户号
	PaySecret string `json:"pay_secret" gorm:"type:varchar(50)"`                 // 商户api秘钥
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
