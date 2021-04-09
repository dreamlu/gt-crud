// author: dreamlu
package order

import (
	"demo/models"
	"demo/util/result"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/type/cmap"
)

// order model
type Order struct {
	models.AdminCom
	ClientID uint64 `json:"client_id" gorm:"type:bigint(20);INDEX:查询索引client_id"` // 客户id
	// 0待付款(取消支付),1待发货,2待收货,3已完成,4退款完成,5申请退款中,6拒绝退款,7待评价
	Status     *int8   `json:"status" gorm:"type:tinyint(2);DEFAULT:0"`
	Money      float64 `json:"money" gorm:"type:decimal(10,2)"`      // 付款金额
	OutTradeNo string  `json:"out_trade_no" gorm:"type:varchar(50)"` // 商户订单号(退款等用)
}

// ========== 以下为 多表连接示例, 不用可删除 ================

// order detail
type OrderD struct {
	Order
	ClientName string `json:"client_name"` // client.name
}

// get order, limit and search
// clientPage 1, everyPage 10 default
func (c *Order) GetMoreBySearch(params cmap.CMap) (datas []*OrderD, pager result.Pager, err error) {
	var crud = gt.NewCrud(
		gt.Inner("order", "client"),
		//gt.Left("order", "service"),
		gt.Model(OrderD{}),
		gt.Data(&datas),
	)

	cd := crud.GetMoreBySearch(params)
	if cd.Error() != nil {
		return nil, pager, cd.Error()
	}
	pager.Pager = cd.Pager()
	return datas, pager, nil
}
