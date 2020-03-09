// @author  dreamlu
package main

import (
	"demo/models"
	"demo/routers"
	"github.com/dreamlu/gt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	// pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":" + gt.Configger().GetString("app.port"))
}

// 数据库模型自动生成初始化
func init() {
	gt.NewDBTool().AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Order{},
	)
}
