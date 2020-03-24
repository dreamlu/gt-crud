package client

import (
	"demo/models/client"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/validator"
	"github.com/dreamlu/gt/tool/xss"
	"github.com/gin-gonic/gin"
	"net/http"
)

var p client.Client

//根据id获得data
func GetByID(u *gin.Context) {
	var (
		res interface{}
	)
	id := u.Query("id")
	data, err := p.GetByID(id)
	if err != nil {
		res = result.CError(err)
	}
	res = result.GetSuccess(data)
	u.JSON(http.StatusOK, res)
}

//data信息分页
func GetBySearch(u *gin.Context) {
	var (
		res interface{}
	)
	_ = u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	datas, pager, err := p.GetBySearch(values)
	if err != nil {
		res = result.CError(err)
	}
	res = result.GetSuccessPager(datas, pager)
	u.JSON(http.StatusOK, res)
}

//data信息删除
func Delete(u *gin.Context) {
	var (
		res interface{}
	)
	id := u.Param("id")
	err := p.Delete(id)
	if err != nil {
		res = result.CError(err)
	}
	res = result.MapDelete
	u.JSON(http.StatusOK, res)
}

//data信息修改
func Update(u *gin.Context) {
	var (
		data client.Client
		res  interface{}
	)
	// json 类型需要匹配
	// 与spring boot不同
	// 不能自动将字符串转成对应类型
	// 严格匹配
	_ = u.ShouldBindJSON(&data)
	// do something

	_, err := p.Create(&data)
	if err != nil {
		res = result.CError(err)
	}
	res = result.MapCreate
	u.JSON(http.StatusOK, res)
}

//新增data信息
func Create(u *gin.Context) {
	var (
		data client.Client
		res  interface{}
	)

	// 自定义日期格式问题
	_ = u.ShouldBindJSON(&data)

	_, err := p.Create(&data)
	if err != nil {
		res = result.CError(err)
	}
	res = result.MapCreate
	u.JSON(http.StatusOK, res)
}

//data信息修改
func UpdateForm(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)

	ss := p.UpdateForm(values)
	u.JSON(http.StatusOK, ss)
}

//新增data信息
func CreateForm(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)                              //html特殊字符转换
	res := validator.Valid(values, client.Client{}) //验证规则
	if res != result.MapValSuccess {
		u.JSON(http.StatusOK, res)
		return
	}

	ss := p.CreateForm(values)
	u.JSON(http.StatusOK, ss)
}
