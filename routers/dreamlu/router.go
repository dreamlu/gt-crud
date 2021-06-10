package dreamlu

import (
	"demo/controllers"
	applet2 "demo/models/admin/applet"
	client2 "demo/models/client"
	"demo/routers"
	"demo/util/db"
	"demo/util/pool"
	"sync"
)

var cls = map[string]controllers.ComController{}

func InitRouter() {
	// 路由定义
	Route(map[string]interface{}{
		// 客户
		"/client": &client2.Client{}, // 取指针, 则实现方式既可以是指针实现也可以是值实现
		// admin
		"/admin/applet": &applet2.Applet{},
	})
}

func Route(route map[string]interface{}) {

	var g sync.WaitGroup
	//var dbm []interface{}
	for k, v := range route {
		t := v // pointer
		cls[k] = controllers.New(v)
		//dbm = append(dbm, v)
		// 加速初始化
		g.Add(1)
		pool.Gsync.Submit(func() {
			defer g.Done()
			db.InitDB(t)
		})
	}
	//db.InitDBRouter(dbm...)
	g.Wait()

	v := routers.V
	{
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
		// 其他接口
		oRoute(v)
	}
}
