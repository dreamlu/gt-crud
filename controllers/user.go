// @author  dreamlu
package controllers

import (
	"demo/models"
	"github.com/dreamlu/go-tool/util/xss"
	"github.com/dreamlu/go-tool/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

var p models.User
//根据id获得用户获取
func GetById(u *gin.Context) {
	id := u.Query("id")
	ss := p.GetByID(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息分页
func GetBySearch(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	ss := p.GetBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func Delete(u *gin.Context) {
	id := u.Param("id")
	ss := p.Delete(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func Update(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)
	val := validator.NewValidator(values) //验证规则
	val.AddRule("name", "用户名","required,len","2-20")
	info := val.CheckInfo()
	if info != nil {
		u.JSON(http.StatusOK, info)
		return
	}

	ss := p.Update(values)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func Create(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	xss.XssMap(values)	//html特殊字符转换

	val := validator.NewValidator(values) //验证规则
	val.AddRule("name", "用户名","required,len","2-20")
	info := val.CheckInfo()
	if info != nil {
		u.JSON(http.StatusOK, info)
		return
	}

	ss := p.Create(values)
	u.JSON(http.StatusOK, ss)
}
