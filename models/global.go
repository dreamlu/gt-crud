package models

import "github.com/dreamlu/gt/tool/type/time"

type IDCom struct {
	ID uint64 `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
}

// 通用模型
type ModelCom struct {
	IDCom
	Createtime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"createtime"` // 创建时间自动生成
}

// 多账号, 如果不需要多账号, 注释掉AdminID即可(ps: 为了简化部署,可直接在util/db/db.go中加入初始化appid applet账号信息)
// 账号关联
type AdminCom struct {
	ModelCom
	AdminID uint64 `json:"admin_id" gorm:"type:bigint(20);INDEX:查询索引admin_id"`
}
