package routers

import (
	"github.com/Dreamlu/deercoder-gin/controllers"
	"github.com/Dreamlu/deercoder-gin/controllers/basic"
	"github.com/Dreamlu/deercoder-gin/util/file"
	"github.com/Dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//router := gin.Default()
	router := gin.New()
	//deercoder.MaxUploadMemory = router.MaxMultipartMemory
	//router.Use(CorsMiddleware())

	// load the casbin model and policy from files, database is also supported.
	//登录失效验证
	//router.Use(CheckLogin())
	//权限中间件
	//e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	//router.Use(authz.NewAuthorizer(e))

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
		//网站基本信息
		v.GET("/basic/getbasic", basic.GetBasicInfo)
		//文件上传
		v.POST("/file/upload", file.UpoadFile)
		//用户
		user := v.Group("/user")
		{
			user.GET("/getbysearch", controllers.GetUserBySearch)
			user.GET("/getbyid", controllers.GetUserById)
			user.DELETE("/deletebyid/:id", controllers.DeleteUserById)
			user.POST("/create", controllers.CreateUser)
			user.PATCH("/update", controllers.UpdateUser)
		}
		//用户账户数据
		usercount := v.Group("/userinfo")
		{
			usercount.GET("/getbysearch", controllers.GetUserinfoBySearch)
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
		if !strings.Contains(path, "login") && !strings.Contains(path, "/captcha/getkey") && !strings.Contains(path, "/captcha/showimage") {
			_, err := c.Cookie("role_id")
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, lib.MapNoToken)
			}
		}
	}
}

/*跨域解决方案*/
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var filterHost = [...]string{"http://localhost.*", "http://192.168.31.173:3000"}
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
