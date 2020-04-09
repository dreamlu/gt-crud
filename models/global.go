package models

import "github.com/dreamlu/gt/tool/type/time"

// 通用模型
type ModelCom struct {
	ID         uint64     `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
	Createtime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"createtime"` // 创建时间自动生成
}

// 账号关联
type AdminCom struct {
	ModelCom
	AdminID uint64 `json:"admin_id" gorm:"type:bigint(20);INDEX:查询索引admin_id"`
}
