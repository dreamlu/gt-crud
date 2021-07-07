package result

import (
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/util/xss"
	"github.com/gin-gonic/gin"
)

func GetToken(u *gin.Context) string {

	tk := u.Request.Header.Get("token")
	if tk == "" {
		tk = ToCMap(u).Pop("token") // bug, request >= 2 set to header to fix
		u.Request.Header.Set("token", tk)
	}
	return tk
}

func ToCMap(u *gin.Context) cmap.CMap {
	err := u.Request.ParseForm()
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	values := cmap.CMap(u.Request.Form) //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	return values
}
