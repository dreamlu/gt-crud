package main

import (
	"demo/routers"
	"github.com/Dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// sql输出方式
	// "sqlErr"仅打印到err sql至log文件,debug输出到控制台
	deercoder.LogMode("debug")
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + deercoder.GetConfigValue("http_port"))
}