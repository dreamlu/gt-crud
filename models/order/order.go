// author: dreamlu
package order

import (
	"demo/models"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
)

// order model
type Order struct {
	models.AdminCom
	ClientID uint64 `json:"client_id" gorm:"type:bigint(20);INDEX:查询索引client_id"` // 客户id
	// 0待付款(取消支付),1待发货,2待收货,3已完成,4退款完成,5申请退款中,6拒绝退款,7待评价
	Status     *int8   `json:"status" gorm:"type:tinyint(2);DEFAULT:0"`
	Money      float64 `json:"money" gorm:"type:double(10,2)"`       // 付款金额
	OutTradeNo string  `json:"out_trade_no" gorm:"type:varchar(50)"` // 商户订单号(退款等用)
	ClientName string  `json:"client_name" gorm:"type:varchar(50)"`  // 姓名
	// ServiceID   int64      `json:"service_id"`
}

// ========== 以下为 多表连接示例, 不用可删除 ================

// service model
type Service struct {
	ID   int64  `gorm:"type:bigint(20) AUTO_INCREMENT;PRIMARY_KEY;" json:"id"`
	Name string `json:"name" gorm:"type:varchar(30)"`
}

// order detail
type OrderD struct {
	Order
	ServiceID   int64  `json:"service_id"`   // service table id
	ServiceName string `json:"service_name"` // service table column `name`
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
