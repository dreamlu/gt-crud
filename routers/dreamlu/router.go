package dreamlu

import (
	"demo/controllers"
	"demo/controllers/client"
	"demo/routers"
)

func InitRouter() {
	v := routers.V
	{
		//用户
		clients := v.Group("/client")
		{
			clients.GET("/search", client.GetBySearch)
			clients.GET("/id", client.GetByID)
			clients.DELETE("/delete/:id", client.Delete)
			clients.POST("/create", client.Create)
			clients.PATCH("/update", client.Update)
			clients.POST("/createForm", client.CreateForm)
			clients.PUT("/updateForm", client.UpdateForm)
		}
		//订单数据
		orders := v.Group("/order")
		{
			orders.GET("/search", controllers.GetOrderBySearch)
		}
	}
}
