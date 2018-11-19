package main

import (
	"demo/routers"
	"github.com/Dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + deercoder.GetConfigValue("http_port"))
}