// @author  dreamlu
package routers

import (
	"demo/controllers/file"
	str2 "demo/util/cons"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/util/str"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Router = SetRouter()

var V = Router.Group("/api/v1")

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//router := gin.Default()
	router := gin.New()
	str.MaxUploadMemory = router.MaxMultipartMemory
	//router.Use(CorsMiddleware())

	// 过滤器
	router.Use(Filter())
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

		// 静态目录
		// relativePath:请求路径
		// root:静态文件所在目录
		v.Static("static", "static")
		// v.GET("/statics/file", file.StaticFile)
		//文件上传
		v.POST("/file/upload", file.UploadFile)
		v.POST("/file/multi_upload", file.UploadMultiFile)
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

// 登录失效验证
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 127请求且本地开发且为dev时无需验证,方便自己测试
		if strings.Contains(c.Request.RemoteAddr, "127.0.0.1") &&
			str2.DevMode == str.Dev {
			return
		}

		if c.Request.Method == "GET" {
			//c.Next()
			return
		}
		path := c.Request.URL.String()

		// 静态服务器 file 处理
		if strings.Contains(path, "/static/file/") {
			file.StaticFile(c)
			c.Abort()
			return
		}

		if !strings.Contains(path, "login") && !strings.Contains(path, "/static/file") {
			_, err := c.Cookie("uid") // may be use session
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, result.MapNoAuth)
				return
			}
		}
	}
}

// 处理跨域请求,支持options访问
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		//fmt.Println(method)
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}
