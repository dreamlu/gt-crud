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
		user := v.Group("/user")
		{
			user.GET("/search", client.GetBySearch)
			user.GET("/id", client.GetByID)
			user.DELETE("/delete/:id", client.Delete)
			user.POST("/create", client.Create)
			user.PATCH("/update", client.Update)
			user.POST("/createForm", client.CreateForm)
			user.PUT("/updateForm", client.UpdateForm)
		}
		//订单数据
		orders := v.Group("/order")
		{
			orders.GET("/search", controllers.GetOrderBySearch)
		}
	}
}
