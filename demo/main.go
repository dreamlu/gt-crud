// author:  dreamlu
package main

import (
	"demo/routers"
	"github.com/dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// sql输出方式
	// "sqlErr"仅打印到err sql至log文件,debug输出到控制台
	deercoder.LogMode("debug")
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	//pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":" + deercoder.GetConfigValue("http_port"))
}