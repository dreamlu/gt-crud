// @author  dreamlu
package routers

import (
	"crypto/md5"
	"demo/controllers/common/captcha"
	"demo/controllers/common/file"
	"demo/controllers/common/qrcode"
	"demo/routers/authz"
	"demo/routers/routelist"
	"demo/routers/whitelist"
	"demo/util"
	str2 "demo/util/cons"
	"demo/util/result"
	"fmt"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/util/cons"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strings"
)

var (
	prefix = str2.Prefix
	V      *routelist.Routes
)

func init() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//router := gin.Default()
	router := gin.New()
	// router.MaxMultipartMemory = // 默认32M
	router.Use(Cors())

	// 过滤器
	router.Use(Filter())
	router.Use(Recovery())

	router.Use(authz.NewAuthorizer(authz.Enforcer))
	routelist.RouteList = routelist.NewRoute(router)
	SetRouter(routelist.RouteList)
	V = routelist.RouteList.Group(prefix)
}

func SetRouter(router *routelist.Routes) {

	//组的路由,version
	v1 := router.Group(prefix)
	{
		v := v1

		// Ping test
		v.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		// 静态目录
		// relativePath:请求路径
		// root:静态文件所在目录
		v.Static("static", "static")
		v.Static("template", "conf/template")
		// v.GET("/statics/file", file.StaticFile)
		//文件上传
		v.POST("/file/upload", file.UploadFile)
		v.POST("/file/multi_upload", file.UploadMultiFile)
		v.GET("/captcha", captcha.Captcha)
		v.GET("/qrcode", qrcode.PQrcode)
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":    "接口不存在->('.')/请求方法类型GET/POST...不正确",
		})
	})
}

// 登录失效验证
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 127请求且本地开发且为dev时无需验证,方便自己测试
		if whitelist.WLIp.Contains(c.Request.RemoteAddr) &&
			str2.DevMode == cons.Dev {
			c.Next()
			return
		}

		//if c.Request.Method == "GET" {
		//	c.Next()
		//	return
		//}
		path := c.Request.URL.String()

		// 静态服务器 file 处理
		if util.Contains(path, "/static/file/") {
			file.StaticFile(c)
			c.Abort()
			return
		}

		if whitelist.WLPath.Contains(path) {
			c.Next()
			return
		}

		token := result.GetToken(c)
		if token == "" {
			c.Abort()
			c.JSON(http.StatusOK, result.GetError("缺少token"))
			return
		}
		ca := cache.NewCache()
		log.Info("[token]:", token)
		cam, err := ca.Get(token)
		if err != nil || cam.Data == nil {
			c.Abort()
			c.JSON(http.StatusOK, result.MapNoAuth)
			return
		}
		// 延长token对应时间
		_ = ca.Set(token, cam)

		// 重复点击
		//switch r.Method {
		//case "POST", "PATCH":
		//	b := check(token, path)
		//	if !b {
		//		c.Abort()
		//		c.JSON(http.StatusOK, result.TextError("点击太频繁"))
		//		return
		//	}
		//}
	}
}

// 重复请求全局验证
func check(token, path string) bool {

	// 白名单
	if b := white(path); b {
		return true
	}
	// 判断重复下单:redis
	key := token + path
	// md5加密缩短长度key
	has := md5.Sum([]byte(key))
	key = strings.ToUpper(fmt.Sprintf("%x", has))

	ce := cache.NewCache()
	ca, _ := ce.Get(key)
	if ca.Data == nil {
		ca.Data = 1
		ca.Time = 2 * cache.CacheSecond
		_ = ce.Set(key, ca)
		//return nil
	} else {
		return false
	}
	return true
}

var whitePath = []string{
	"/client/follow",
	"/client/play",
}

// 白名单
func white(path string) bool {
	for _, v := range whitePath {
		if strings.Contains(path, v) {
			return true
		}
	}
	return false
}

// 异常捕获
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			//log.Error(string(debug.Stack()))
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			ss := strings.Split(string(debug.Stack()), "\n\t")
			res := make(map[string]string)
			res["error"] = err
			for _, v := range ss {
				ks := strings.Split(v, "\n")
				res[ks[0]] = ks[1]
			}
			c.JSON(http.StatusOK, result.GetError(res))
		}
		c.Abort()
	})
}

func RecoverOld(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			//debug.PrintStack()
			log.Error(string(debug.Stack()))
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			ss := strings.Split(string(debug.Stack()), "\n\t")
			res := make(map[string]string)
			for _, v := range ss {
				ks := strings.Split(v, "\n")
				res[ks[0]] = ks[1]
			}
			c.JSON(http.StatusOK, result.GetError(res))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
