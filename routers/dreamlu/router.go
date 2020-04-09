package dreamlu

import (
	"demo/controllers"
	"demo/controllers/admin"
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
			admins.GET("/id", admin.Search)
			admins.DELETE("/delete/:id", admin.Delete)
			admins.POST("/create", admin.Create)
			admins.PUT("/update", admin.Update)
			admins.POST("/login", admin.Login)
		}
		//订单数据
		orders := v.Group("/order")
		{
			orders.GET("/search", controllers.GetOrderBySearch)
		}
	}
}
