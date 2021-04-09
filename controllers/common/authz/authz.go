// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package authz

import (
	"demo/util/result"
	"github.com/dreamlu/gt/tool/util"
	"net/http"
	"net/url"
	"strings"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}

		if !a.CheckPermission(c.Request) {
			c.Abort()
			c.JSON(http.StatusOK, result.MapNoAuth)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the data name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	//dataname, _, _ := r.BasicAuth()
	cookie, err := r.Cookie("role_id")
	if err != nil {
		return "-1"
	}
	ss, _ := url.QueryUnescape(cookie.Value)
	// 解密
	role_id := util.AesDe(ss)
	//if err != nil {
	//	fmt.Println("cookie解密失败: ", err)
	//	return "-1"
	//}
	//各位数角色职位,截取一位即可
	return string(role_id[:1])
}

// CheckPermission checks the data/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	data := a.GetUserName(r)
	method := r.Method
	path := r.URL.Path
	if strings.Contains(path, "/static/") || strings.Contains(path, "login") {
		return true
	}
	return a.enforcer.Enforce(data, path, method)
}

/*// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
}*/
