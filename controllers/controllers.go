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
	if f, ok := c.M().(models.GetService); ok {
		u.JSON(http.StatusOK, result.ResGet(f.Get(result.ToCMap(u))))
		return
	}
	u.JSON(http.StatusOK, result.ResGet(c.Service.Get(result.ToCMap(u))))
}

//data信息分页
func (c ComController) Search(u *gin.Context) {
	if f, ok := c.M().(models.SearchService); ok {
		u.JSON(http.StatusOK, result.ResPager(f.Search(result.ToCMap(u))))
		return
	}
	u.JSON(http.StatusOK, result.ResPager(c.Service.Search(result.ToCMap(u))))
}

//data信息删除
func (c ComController) Delete(u *gin.Context) {
	id := u.Param("id")
	if f, ok := c.M().(models.DeleteService); ok {
		u.JSON(http.StatusOK, result.Res(f.Delete(id)))
		return
	}
	u.JSON(http.StatusOK, result.Res(c.Service.Delete(id)))
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
	if f, ok := c.M().(models.UpdateService); ok {
		u.JSON(http.StatusOK, result.Res(f.Update(data)))
		return
	}
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
	if f, ok := c.M().(models.CreateService); ok {
		u.JSON(http.StatusOK, result.Res(f.Create(data)))
		return
	}
	u.JSON(http.StatusOK, result.Res(c.Service.Create(data)))
}
