// @author  dreamlu
package main

import (
	"demo/routers"
	"demo/routers/dreamlu"
	"demo/util/db"
	"github.com/dreamlu/gt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//r := routers.SetRouter()
	// pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = routers.Router.Run(":" + gt.Configger().GetString("app.port"))
}

// 数据库模型自动生成
func init() {
	dreamlu.InitRouter()
	db.InitDB()
}
