package controllers

import (
	"demo/models"
	"demo/util/cm"
	"demo/util/reflect"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ComController struct {
	Com models.Com
}

func New(model interface{}, arrayModel interface{}) ComController {
	return ComController{Com: models.Com{
		Model:      model,
		ArrayModel: arrayModel,
	}}
}

//根据id获得data
func (c *ComController) Get(u *gin.Context) {
	data, err := c.Com.Get(cm.ToCMap(u))
	u.JSON(http.StatusOK, cm.ResGet(err, data))
}

//data信息分页
func (c *ComController) Search(u *gin.Context) {
	datas, pager, err := c.Com.Search(cm.ToCMap(u))
	u.JSON(http.StatusOK, cm.ResPager(err, datas, pager))
}

//data信息删除
func (c *ComController) Delete(u *gin.Context) {
	id := u.Param("id")
	err := c.Com.Delete(id)
	u.JSON(http.StatusOK, cm.Res(err))
}

//data信息修改
func (c *ComController) Update(u *gin.Context) {
	var (
		data = reflect.New(c.Com.Model)
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

	err = c.Com.Update(data)
	u.JSON(http.StatusOK, cm.Res(err))
}

//新增data信息
func (c *ComController) Create(u *gin.Context) {
	var (
		data = reflect.New(c.Com.Model)
	)
	// 自定义日期格式问题
	err := u.ShouldBindJSON(data)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	err = c.Com.Create(data)
	u.JSON(http.StatusOK, cm.Res(err))
}
