package dreamlu

import (
	"demo/controllers/admin"
	"demo/controllers/client"
	"demo/controllers/order"
	"demo/routers/routelist"
)

// ==== 额外接口单独定义 =======
func oRoute(v *routelist.Routes) {
	// 初始化模块路由
	WxRouter(v)

	// 用户-额外接口
	clients := v.Group("/client")
	{
		clients.GET("/token", client.Token)
	}
	// admin
	admins := v.Group("/admin")
	{
		admins.GET("/search", admin.Search)
		admins.GET("/get", admin.Get)
		admins.DELETE("/delete/:id", admin.Delete)
		admins.POST("/create", admin.Create)
		admins.PUT("/update", admin.Update)
		admins.POST("/login", admin.Login)
	}
	//订单数据
	orders := v.Group("/order")
	{
		orders.GET("/search", order.GetOrderBySearch)
	}
}
