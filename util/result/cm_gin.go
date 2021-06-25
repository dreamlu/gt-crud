package result

import (
	"encoding/json"
	"github.com/dreamlu/gt/tool/reflect"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func GinResult(u *gin.Context, result Resultable) {
	u.JSON(http.StatusOK, result)
}

func GinGet(u *gin.Context, f func(params cmap.CMap) (data interface{}, err error), result ...Resultable) {
	GinResult(u, ResGet(f(ToCMap(u))).AddStruct(St(result...)))
}

func GinSearch(u *gin.Context, f func(params cmap.CMap) (datas interface{}, pager Pager, err error), result ...Resultable) {
	GinResult(u, ResPager(f(ToCMap(u))).AddStruct(St(result...)))
}

func GinDelete(u *gin.Context, f func(id interface{}) error, result ...Resultable) {
	GinResult(u, Res(f(u.Param("id"))).AddStruct(St(result...)))
}

func GinCu(u *gin.Context, data interface{}, f func(interface{}) error, result ...Resultable) {
	err := f(data)
	if err != nil {
		GinResult(u, Res(f(data)))
		return
	}
	GinResult(u, Res(nil).AddStruct(St(result...)))
}

func GinCrUp(u *gin.Context, data interface{}, f func(id interface{}) error, result ...Resultable) {
	err := u.ShouldBindJSON(data)
	if err != nil {
		GinResult(u, CError(err))
		return
	}
	GinCu(u, data, f, result...)
}

// GinCreate Create/CreateMore
func GinCreate(u *gin.Context, model interface{}, f func(interface{}) error) {

	var (
		data = reflect.New(model)
		//err  error
	)
	b, err := ioutil.ReadAll(u.Request.Body)
	if err != nil {
		GinResult(u, CError(err))
		return
	}
	err = json.Unmarshal(b, data)
	//decoder := json.NewDecoder(u.Request.Body)
	//err = decoder.Decode(data)
	if err != nil {
		// 后续优化判断方式, 字段属性可能会有影响
		if strings.Contains(err.Error(), "cannot unmarshal array into Go value") {
			err = nil
			data = reflect.NewArray(model)
			err = json.Unmarshal(b, data)
			if err != nil {
				GinResult(u, CError(err))
				return
			}
			GinCu(u, data, f)
			return
		}
		GinResult(u, CError(err))
		return
	}

	GinCu(u, data, f, Add("data", data))
}

func St(result ...Resultable) Resultable {
	if len(result) == 1 {
		return result[0]
	}
	return nil
}
