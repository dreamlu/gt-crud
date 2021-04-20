package dreamlu

import (
	"demo/controllers"
	"demo/controllers/admin"
	"demo/controllers/client"
	"demo/controllers/order"
	applet2 "demo/models/admin/applet"
	client2 "demo/models/client"
	"demo/routers"
	"demo/util/db"
	"github.com/gin-gonic/gin"
)

var cls = map[string]controllers.ComController{}

func InitRouter() {
	v := routers.V
	{
		// 路由定义
		Route(map[string]interface{}{
			// 客户
			"/client": &client2.Client{}, // 取指针, 则实现方式既可以是指针实现也可以是值实现
			// admin
			"/admin/applet": &applet2.Applet{},
		})

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

		oRoute(v)
	}
}

func Route(route map[string]interface{}) {

	var dbm []interface{}
	for k, v := range route {
		cls[k] = controllers.New(v)
		dbm = append(dbm, v)
	}
	db.InitDB(dbm...)
}

// ==== 额外接口单独定义 =======
func oRoute(v *gin.RouterGroup) {
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
	// 初始化模块路由
	WxRouter()
	//订单数据
	orders := v.Group("/order")
	{
		orders.GET("/search", order.GetOrderBySearch)
	}
}
