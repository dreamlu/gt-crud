package dreamlu

import (
	"demo/controllers"
	"demo/controllers/admin"
	"demo/controllers/admin/applet"
	"demo/controllers/client"
	"demo/routers"
)

func InitRouter() {
	v := routers.V
	{
		//用户
		clients := v.Group("/client")
		{
			clients.GET("/search", client.Search)
			clients.GET("/id", client.Get)
			clients.DELETE("/delete/:id", client.Delete)
			clients.POST("/create", client.Create)
			clients.PATCH("/update", client.Update)
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
			orders.GET("/search", controllers.GetOrderBySearch)
		}
	}
}
