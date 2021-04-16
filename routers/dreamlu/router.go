package dreamlu

import (
	"demo/controllers"
	"demo/controllers/admin"
	"demo/controllers/admin/applet"
	"demo/controllers/client"
	"demo/controllers/order"
	applet2 "demo/models/admin/applet"
	client2 "demo/models/client"
	"demo/routers"
)

var cls = map[string]controllers.ComController{}

func InitRouter() {
	v := routers.V
	{
		// 用户
		cls["/client"] = controllers.New(client2.Client{})

		// 小程序账号
		cls["/admin/applet"] = controllers.New(applet2.Applet{})

		// 路由列表
		for k, c := range cls {
			pre := v.Group(k)
			{
				pre.GET("/search", c.Search)
				pre.GET("/get", c.Get)
				pre.DELETE("/delete/:id", c.Delete)
				pre.POST("/create", c.Create)
				pre.PUT("/update", c.Update)
			}
		}

		// ==== 额外接口单独定义 =======
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

			// 小程序账号
			applets := admins.Group("/applet")
			{
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
