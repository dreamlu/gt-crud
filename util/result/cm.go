package result

import (
	"errors"
	"github.com/dreamlu/gt/tool/log"
	"github.com/dreamlu/gt/tool/type/cmap"
	te "github.com/dreamlu/gt/tool/type/errors"
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

func Res(err error) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = MapSuccess
	}
	return
}

func ResGet(err error, data interface{}) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = GetSuccess(data)
	}
	return
}

func ResPager(err error, datas interface{}, pager Pager) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		if pager.TotalNum == 0 {
			datas = []interface{}{}
		}
		res = GetSuccessPager(datas, pager)
	}
	return
}
