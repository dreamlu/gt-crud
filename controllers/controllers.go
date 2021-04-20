package controllers

import (
	"demo/models"
	"demo/util/result"
	"github.com/dreamlu/gt/tool/reflect"
	"github.com/gin-gonic/gin"
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

func New(model interface{}) ComController {
	return ComController{Service: models.NewService(model)}
}

// Get 获得data
func (c ComController) Get(u *gin.Context) {
	if f, ok := c.M().(models.GetService); ok {
		result.GinGet(u, f.Get)
		return
	}
	result.GinGet(u, c.Service.Get)
}

// Search data信息分页
func (c ComController) Search(u *gin.Context) {
	if f, ok := c.M().(models.SearchService); ok {
		result.GinSearch(u, f.Search)
		return
	}
	result.GinSearch(u, c.Service.Search)
}

// Delete data信息删除
func (c ComController) Delete(u *gin.Context) {
	if f, ok := c.M().(models.DeleteService); ok {
		result.GinDelete(u, f.Delete)
		return
	}
	result.GinDelete(u, c.Service.Delete)
}

// Update data信息修改
func (c ComController) Update(u *gin.Context) {
	var (
		data = reflect.New(c.Service.M())
	)
	// json 类型需要匹配
	// 严格匹配
	if f, ok := c.M().(models.UpdateService); ok {
		result.GinCrUp(u, data, f.Update)
		return
	}
	result.GinCrUp(u, data, c.Service.Update)
}

// Create 新增data信息
func (c ComController) Create(u *gin.Context) {
	var (
		data = reflect.New(c.Service.M())
	)
	if f, ok := c.M().(models.CreateService); ok {
		result.GinCrUp(u, data, f.Create, result.Add("data", data))
		return
	}
	result.GinCrUp(u, data, c.Service.Create, result.Add("data", data))
}
