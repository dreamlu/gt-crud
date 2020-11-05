package wx

import (
	"demo/models/admin/applet"
	"demo/models/order"
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/id"
	log2 "github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/payment"
	"log"
	"net/http"
)

type WxOrder struct {
	OrderID uint64 `json:"order_id"` // 订单id
	applet.Applet
}

// 退款
func Refund(u *gin.Context) {
	var (
		wx WxOrder
		or order.Order
	)
	_ = u.ShouldBindJSON(&wx)

	// order参数查询
	//err := or.GetByID(wx.OrderID)
	//if err != nil {
	//	u.JSON(http.StatusOK, result.GetError(err.Error()))
	//	return
	//}

	// applet参数查询
	//if err := wx.GetByAdminID(or.AdminID); err != nil {
	//	u.JSON(http.StatusOK, result.GetError(err.Error()))
	//	return
	//}

	refundNo, _ := id.NewID(1)
	notifyUrl := conf.GetString("app.notifyUrl") + "/refund"
	// 新建退款订单
	form := payment.Refunder{
		// 必填
		AppID:       wx.Appid,
		MchID:       wx.MchID,
		TotalFee:    int(or.Money * 100), //"总金额(分)"
		RefundFee:   int(or.Money * 100), //"退款金额(分)"
		OutRefundNo: refundNo.String(),
		// 二选一
		OutTradeNo: or.OutTradeNo, // or TransactionID: "微信订单号",
		// 选填 ...
		RefundDesc: "用户退款",    // 若商户传入, 会在下发给用户的退款消息中体现退款原因
		NotifyURL:  notifyUrl, //结果通知地址，覆盖商户平台上配置的回调地址
	}

	// 需要证书
	res, err := form.Refund(wx.PaySecret, wx.AppCert, wx.AppKey)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	log.Printf("返回结果: %#v", res)
	//err = o.EditStatus(id, 4)
	//if err != nil {
	//	u.JSON(http.StatusOK, result.GetMapData(500, "退款失败"))
	//	return
	//}
	u.JSON(http.StatusOK, res.OutRefundNo)
}

func RefundNotify(u *gin.Context) {
	var wx applet.Applet
	// applet参数查询
	if err := wx.GetByAppid(wx.Appid); err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	// 必须在商户平台上配置的回调地址或者发起退款时指定的 notify_url 的路由处理器下
	w := u.Writer
	req := u.Request
	err := payment.HandleRefundedNotify(w, req, wx.PaySecret, func(notify payment.RefundedNotify) (b bool, s string) {
		//OutTradeNo := notify.OutTradeNo
		status := notify.RefundStatus
		if status == "SUCCESS" {
			// 订单状态修改
			//s := int8(4)
			//var or = &order.Order{
			//	OutTradeNo: OutTradeNo,
			//	Status:     &s,
			//}
			//_, err := or.Update(or)
			//if err != nil {
			//	return false, "失败原因..." + err.Error()
			//}
			return true, ""
		}
		var msg string
		if status == "CHANGE" {
			msg = "退款异常"
		} else {
			msg = status
		}
		log2.Error(msg)
		return false, msg
	})
	if err != nil {
		log2.Error(err)
	}
}
