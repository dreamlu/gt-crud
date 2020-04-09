package db

import (
	"demo/models"
	"demo/models/admin"
	"demo/models/client"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/util"
)

func InitDB() {
	gt.NewDBTool().AutoMigrate(
		&client.Client{},
		&models.Service{},
		&models.Order{},
		&admin.Admin{},
	)
	initSQL()
}

// 初次部署更新, 后续可注释取消掉
func initSQL() {
	// 插入admin账号
	role := int8(1)
	var ad = admin.Admin{
		Name:     "admin",
		Password: util.AesEn("123456"),
		Role:     &role,
	}
	gt.NewCrud(gt.Data(&ad), gt.Model(admin.Admin{})).Create()
}
