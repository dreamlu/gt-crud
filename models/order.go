// author: dreamlu
package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/time"
)

// order model
type Order struct {
	ID         int64      `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
	UserID     int64      `json:"user_id" gorm:"type:bigint(20)"`    // user id
	ServiceID  int64      `json:"service_id" gorm:"type:bigint(20)"` // service table id
	CreateTime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"createtime"`
}

// service model
type Service struct {
	ID   int64  `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
	Name string `json:"name" gorm:"type:varchar(30)"`
}

// order detail
type OrderD struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`      // user id
	UserName    string     `json:"user_name"`    // user table column name
	ServiceID   int64      `json:"service_id"`   // service table id
	ServiceName string     `json:"service_name"` // service table column `name`
	Createtime  time.CTime `json:"createtime"`   // createtime
}

// get order, limit and search
// clientPage 1, everyPage 10 default
func (c *Order) GetMoreBySearch(params map[string][]string) interface{} {
	var or []OrderD
	var crud = gt.NewCrud(
		gt.InnerTable([]string{"order", "user"}),
		gt.LeftTable([]string{"order", "service"}),
		gt.Model(OrderD{}),
		gt.Data(&or),
	)

	cd := crud.GetMoreBySearch(params)
	if cd.Error() != nil {
		return result.GetError(cd.Error())
	}

	return result.GetSuccessPager(or, cd.Pager())
}
