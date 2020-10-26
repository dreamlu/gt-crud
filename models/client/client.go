package client

import (
	"demo/models"
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

func (c Client) New() models.DN {
	return Client{}
}
