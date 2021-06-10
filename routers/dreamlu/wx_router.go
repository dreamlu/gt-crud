package dreamlu

import (
	"demo/controllers/wx"
	"demo/routers/routelist"
)

func WxRouter(v *routelist.Routes) {
	// 小程序
	wxs := v.Group("/wx")
	{
		wxs.POST("/login", wx.Login)
		wxs.GET("/info", wx.Info)
		wxs.GET("/phone", wx.Phone)
		wxs.POST("/pay", wx.Pay)
		wxs.GET("/access_token", wx.GetAccessToken)
		wxs.GET("/qrcode", wx.GetQRCode)
		wxs.GET("/qrcode/key", wx.GetByKey)
		wxs.POST("/refund", wx.Refund)

		// 回调
		wxs.POST("/notify/pay", wx.PayNotify)
		wxs.POST("/notify/refund", wx.RefundNotify)
	}
}
