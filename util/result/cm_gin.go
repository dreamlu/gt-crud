package result

import (
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/gin-gonic/gin"
	"net/http"
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

func GinCrUp(u *gin.Context, data interface{}, f func(id interface{}) error, result ...Resultable) {
	err := u.ShouldBindJSON(data)
	if err != nil {
		GinResult(u, CError(err))
		return
	}
	GinResult(u, Res(f(data)).AddStruct(St(result...)))
}

func St(result ...Resultable) Resultable {
	if len(result) == 1 {
		return result[0]
	}
	return nil
}
