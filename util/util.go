package util

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/xss"
	"github.com/gin-gonic/gin"
)

func ToCMap(u *gin.Context) cmap.CMap {
	err := u.Request.ParseForm()
	if err != nil {
		gt.Logger().Error(err.Error())
		return nil
	}
	values := cmap.CMap(u.Request.Form) //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	return values
}
