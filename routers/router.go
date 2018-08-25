package routers

import (
	"deercoder-gin/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"deercoder-gin/util/file"
)

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(CorsMiddleware())

	//静态目录
	router.Static("static", "../static")

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//组的路由,version
	v1 := router.Group("/api/v1")
	{
		v := v1
		//文件上传
		v.POST("/file/upload",file.UpoadFile)
		//用户
		user := v.Group("/user")
		{
			user.GET("/getbysearch", controllers.GetUserBySearch)
			user.GET("/getbyid", controllers.GetUserById)
			user.DELETE("/deletebyid/:id", controllers.DeleteUserById)
			user.POST("/create", controllers.CreateUser)
			user.PATCH("/update", controllers.UpdateUser)
		}
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":  "接口不存在->('.')",
		})
	})
	return router
}

/*跨域解决方案*/
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var filterHost = [...]string{"http://localhost.*", "http://192.168.31.174:3000"}
		// filterHost 做过滤器，防止不合法的域名访问
		var isAccess = false
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
