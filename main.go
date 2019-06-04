// @author  dreamlu
package main

import (
	"demo/routers"
	"github.com/dreamlu/go-tool"
	"github.com/gin-gonic/gin"
)

func main() {
	// open db
	der.NewDB()
	// sql输出方式
	// "sqlErr"仅打印到err sql至log文件,debug输出到控制台
	der.LogMode("debug")
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	//pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":" + der.GetConfigValue("http_port"))
}