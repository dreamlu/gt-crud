package db

import (
	"demo/controllers/wx"
	"demo/models/admin"
	"demo/models/admin/applet"
	"demo/models/client"
	"demo/models/order"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/util"
)

func InitDB() {
	gt.DB().AutoMigrate(
		&client.Client{}, // 客户
		&order.Order{},   // 订单, 不用可注释
		&admin.Admin{},   // 账号管理
		&applet.Applet{}, // appid账号存储
		&wx.QrCode{},     // 小程序二维码解析参数存储
	)
	initSQL()
}

// 初次部署更新, 后续可注释取消掉
func initSQL() {
	initAdmin()
	//initApplet()
}

// 单账号初始化
// 多账号可不用
func initApplet() {
	var data = applet.Applet{
		//AdminCom:  models.AdminCom{}, // 单账号不需要admin_id,注释掉
		Appid:     "",
		Secret:    "",
		MchID:     "",
		PaySecret: "",
		AppCert:   "",
		AppKey:    "",
	}
	if data.Appid == "" {
		return
	}
	gt.NewCrud(gt.Data(&data), gt.Model(applet.Applet{})).Create()
}

// init admin
func initAdmin() {
	// 插入admin账号
	role := int8(0)
	var ad = admin.Admin{
		Name:     "admin",
		Password: util.AesEn("123456"),
		Role:     &role,
	}
	gt.NewCrud(gt.Data(&ad), gt.Model(admin.Admin{})).Create()
}
