// @author  dreamlu
package main

import (
	"demo/routers"
	"demo/util/global"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	// pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":" + global.Config.GetString("app.port"))
}
