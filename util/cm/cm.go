package cm

import (
	"errors"
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/type/te"
	"github.com/dreamlu/gt/tool/util/xss"
	"github.com/gin-gonic/gin"
)

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

func Res(err error) (res result.Resultable) {
	if err != nil {
		res = result.CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = result.MapSuccess
	}
	return
}

func ResGet(err error, data interface{}) (res result.Resultable) {
	if err != nil {
		res = result.CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = result.GetSuccess(data)
	}
	return
}

func ResPager(err error, datas interface{}, pager result.Pager) (res result.Resultable) {
	if err != nil {
		res = result.CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = result.GetSuccessPager(datas, pager)
	}
	return
}
