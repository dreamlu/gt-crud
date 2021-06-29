package authz

import (
	"demo/routers/whitelist"
	str2 "demo/util/cons"
	"demo/util/result"
	"demo/util/token"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/file/file_func"
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/util/cons"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	Enforcer *casbin.Enforcer
)

func init() {
	//权限中间件
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	ds := fmt.Sprintf("root:%s@%s/", conf.GetString("app.db.password"), conf.GetString("app.db.host"))
	var err error
	adapter, err := gormadapter.NewAdapter("mysql", ds, conf.GetString("app.db.name"))
	if err != nil {
		panic(err.Error())
	}
	path := file_func.ProjectPath()
	if path != "" {
		path += "/"
	}
	Enforcer, err = casbin.NewEnforcer(path+"conf/authz_model.conf", adapter)
	if err != nil {
		panic(err.Error())
	}
	// Load the policy from DB.
	err = Enforcer.LoadPolicy()
	if err != nil {
		log.Error(err)
		panic(err.Error())
	}
}

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		if whitelist.WLIp.Contains(c.Request.RemoteAddr) &&
			str2.DevMode == cons.Dev {
			c.Next()
			return
		}

		if whitelist.WLPath.Contains(c.Request.URL.String()) {
			c.Next()
			return
		}

		a := NewBasicAuthorizer(e)
		if !a.CheckPermission(c.Request) {
			c.Abort()
			c.JSON(http.StatusOK, result.MapNoAuth)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
	redis    cache.Cache
}

func NewBasicAuthorizer(e *casbin.Enforcer) *BasicAuthorizer {
	return &BasicAuthorizer{enforcer: e, redis: cache.NewCache()}
}

// GetRole 获得角色
func (a *BasicAuthorizer) GetRole(r *http.Request) string {
	tk := r.Header.Get("token")
	role, g, err := token.GetRole(tk)
	if err != nil {
		return "no role"
	}
	// 组存在则使用组
	if g != "" {
		return g
	}
	// 组不存在则必定是一个角色
	if len(role) > 0 {
		return role[0]
	}
	return ""
}

// CheckPermission checks the data/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	role := a.GetRole(r)
	// 所有权限
	//if role == str2.RoleAdmin {
	//	return true
	//}
	method := r.Method
	path := strings.TrimPrefix(r.URL.Path, str2.Prefix)
	b, err := a.enforcer.Enforce(role, path, method)
	return b && err == nil
}
