package dreamlu

import (
	"demo/controllers"
	"demo/controllers/admin"
	"demo/controllers/admin/applet"
	"demo/controllers/client"
	"demo/controllers/order"
	client2 "demo/models/client"
	"demo/routers"
)

func InitRouter() {
	v := routers.V
	{
		//用户
		clients := v.Group("/client")
		{
			clientCon := controllers.New(client2.Client{}, []*client2.Client{})
			clients.GET("/token", client.Token)
			clients.GET("/search", clientCon.Search)
			clients.GET("/id", clientCon.Get)
			clients.DELETE("/delete/:id", clientCon.Delete)
			clients.POST("/create", clientCon.Create)
			clients.PUT("/update", clientCon.Update)
		}
		// admin
		admins := v.Group("/admin")
		{
			admins.GET("/search", admin.Search)
			admins.GET("/id", admin.Get)
			admins.DELETE("/delete/:id", admin.Delete)
			admins.POST("/create", admin.Create)
			admins.PUT("/update", admin.Update)
			admins.POST("/login", admin.Login)

			// 小程序账号
			applets := admins.Group("/applet")
			{
				applets.GET("/search", applet.Search)
				applets.GET("/id", applet.Get)
				applets.DELETE("/delete/:id", applet.Delete)
				applets.POST("/create", applet.Create)
				applets.PUT("/update", applet.Update)
				applets.POST("/download", applet.DownLoad)
			}
		}
		// 初始化模块路由
		WxRouter()
		//订单数据
		orders := v.Group("/order")
		{
			orders.GET("/search", order.GetOrderBySearch)
		}
	}
}
