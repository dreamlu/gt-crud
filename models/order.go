// author: dreamlu
package models

import (
	. "demo/util/global"
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/tool/result"
	"github.com/dreamlu/go-tool/tool/type/time"
)

// order
type Order struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`     // user id
	ServiceID  int64 `json:"service_id"`  // service table id
	CreateTime int64 `json:"create_time"` // createtime
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
	DBTool.Param = der.CrudParam{
		InnerTables: []string{"order", "user"}, // inner join tables, 'order' must the first table
		LeftTables:  []string{"service"},       // left join tables
		Model:       OrderD{},                  // order model
		ModelData:   &or,                       // model value
	}
	pager, err := DBTool.Crud.GetMoreBySearch(params)
	if err != nil {
		return result.GetError(err)
	}

	return result.GetSuccessPager(or, pager)
}
