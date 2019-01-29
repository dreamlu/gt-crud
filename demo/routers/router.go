// author:  dreamlu
package routers

import (
	"demo/controllers"
	"demo/controllers/basic"
	"github.com/dreamlu/deercoder-gin/util/file"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//router := gin.Default()
	router := gin.New()
	//deercoder.MaxUploadMemory = router.MaxMultipartMemory
	//router.Use(CorsMiddleware())

	//登录失效验证
	//router.Use(CheckLogin())
	//权限中间件
	// load the casbin model and policy from files, database is also supported.
	//e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	//router.Use(authz.NewAuthorizer(e))

	//cookie session
	//store := cookie.NewStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	//redis session
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//组的路由,version
	v1 := router.Group("/api/v1")
	{
		v := v1
		//网站基本信息
		v.GET("/basic/basic", basic.GetBasicInfo)
		//文件上传
		v.POST("/file/upload", file.UpoadFile)
		//用户
		user := v.Group("/user")
		{
			user.GET("/search", controllers.GetBySearch)
			user.GET("/id", controllers.GetById)
			user.DELETE("/delete/:id", controllers.Delete)
			user.POST("/create", controllers.Create)
			user.PATCH("/update", controllers.Update)
		}
		//用户账户数据
		usercount := v.Group("/userinfo")
		{
			usercount.GET("/search", controllers.GetUserinfoBySearch)
		}
		//订单数据
		orders := v.Group("/order")
		{
			orders.GET("/search", controllers.GetOrderBySearch)
		}
	}
	//组的路由,version
	v2 := router.Group("/api/v2")
	{
		v := v2

		// 静态目录
		// relativePath:请求路径
		// root:静态文件所在目录
		v.Static("static", "static")

		//网站基本信息
		v.GET("/basic/basic", basic.GetBasicInfo)
		//文件上传
		v.POST("/file/upload", file.UpoadFile)
		//用户
		user := v.Group("/user")
		{
			user.GET("/search", controllers.GetBySearch)
			user.GET("/id", controllers.GetById)
			user.DELETE("/delete/:id", controllers.Delete)
			user.POST("/create", controllers.Create)
			user.PATCH("/update", controllers.Update)
		}
		//用户账户数据
		usercount := v.Group("/userinfo")
		{
			usercount.GET("/search", controllers.GetUserinfoBySearch)
		}
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":    "接口不存在->('.')/请求方法不存在",
		})
	})
	return router
}

/*登录失效验证*/
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.String()
		if !strings.Contains(path, "login") && !strings.Contains(path, "/static/file") {
			_, err := c.Cookie("uid")	// may be use session
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, lib.MapNoAuth)
			}
		}
	}
}

// xss
//func xssMid() gin.HandlerFunc{
//	return func(c *gin.Context) {
//		c.Request.ParseForm()
//		//values := c.Request.PostForm
//		xss.XssMap(c.Request.PostForm)
//	}
//}

///*跨域解决方案,待完善,建议nginx 解决*/
//func CorsMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		origin := c.Request.Header.Get("Origin")
//		var filterHost = [...]string{"http://localhost.*", "http://192.168.31.173:3000"}
//		// filterHost 做过滤器，防止不合法的域名访问
//		var isAccess = false
//		for _, v := range filterHost {
//			match, _ := regexp.MatchString(v, origin)
//			if match {
//				isAccess = true
//			}
//		}
//		if isAccess {
//			// 核心处理方式
//			c.Header("Access-Control-Allow-Origin", "*")
//			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
//			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
//			c.Set("content-type", "application/json")
//		}
//		//放行所有OPTIONS方法
//		if method == "OPTIONS" {
//			c.JSON(http.StatusOK, "Options Request!")
//		}
//		c.Next()
//	}
//}
