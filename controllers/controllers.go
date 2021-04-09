package controllers

import (
	"demo/models"
	"demo/util/result"
	"github.com/dreamlu/gt/tool/reflect"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Controller interface {
//	Get(u *gin.Context)
//	Search(u *gin.Context)
//	Delete(u *gin.Context)
//	Update(u *gin.Context)
//	Create(u *gin.Context)
//}

type ComController struct {
	models.Service
}

func New(model interface{}, params ...models.CrudServiceParam) ComController {
	return ComController{Service: models.NewService(model, params...)}
}

//根据id获得data
func (c ComController) Get(u *gin.Context) {
	data, err := c.Service.Get(result.ToCMap(u))
	u.JSON(http.StatusOK, result.ResGet(err, data))
}

//data信息分页
func (c ComController) Search(u *gin.Context) {
	datas, pager, err := c.Service.Search(result.ToCMap(u))
	u.JSON(http.StatusOK, result.ResPager(err, datas, pager))
}

//data信息删除
func (c ComController) Delete(u *gin.Context) {
	id := u.Param("id")
	err := c.Service.Delete(id)
	u.JSON(http.StatusOK, result.Res(err))
}

//data信息修改
func (c ComController) Update(u *gin.Context) {
	var (
		data = reflect.New(c.Service.M())
	)
	// json 类型需要匹配
	// 与spring boot不同
	// 不能自动将字符串转成对应类型
	// 严格匹配
	err := u.ShouldBindJSON(data)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	// do something

	err = c.Service.Update(data)
	u.JSON(http.StatusOK, result.Res(err))
}

//新增data信息
func (c ComController) Create(u *gin.Context) {
	var (
		data = reflect.New(c.Service.M())
	)
	// 自定义日期格式问题
	err := u.ShouldBindJSON(data)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	err = c.Service.Create(data)
	u.JSON(http.StatusOK, result.Res(err))
}
