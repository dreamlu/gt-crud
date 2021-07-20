// @author  dreamlu
package main

import (
	"demo/routers/dreamlu"
	"demo/routers/routelist"
	"demo/util/cron"
	db2 "demo/util/init/db"
	"demo/util/init/policy"
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/util/cons"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	if conf.GetString("app.devMode") == cons.Dev {
		gin.SetMode(gin.DebugMode)
	} else {
		log.DefaultFileLog()
	}
	//r := routers.SetRouter()
	// 性能调试
	//pprof.Register(routers.Router)
	// Listen and Server in 0.0.0.0:8080
	_ = routelist.RouteList.Router.Run(":" + conf.GetString("app.port"))
}

// 数据库模型自动生成
func init() {
	dreamlu.InitRouter()
	db2.InitDBOther()
	// 如果需要权限,则打开注释
	// 必须路由初始化完成后在初始化权限
	go policy.InitPolicy()
	cron.Cron()
}
