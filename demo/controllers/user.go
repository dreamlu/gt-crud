package controllers

import (
	"demo/models"
	"github.com/Dreamlu/deercoder-gin/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

var user models.User
//根据id获得用户获取
func GetUserById(u *gin.Context) {
	id := u.Query("id")
	ss := user.GetById(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息分页
func GetUserBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := user.GetBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func DeleteUserById(u *gin.Context) {
	id := u.Param("id")
	ss := user.DeleteByid(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func UpdateUser(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	val := validator.NewValidator(values) //验证规则
	val.AddRule("name", "用户名","required,len","2-20")
	info := val.CheckInfo()
	if info != nil {
		u.JSON(http.StatusOK, info)
		return
	}

	ss := user.Update(values)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func CreateUser(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法

	val := validator.NewValidator(values) //验证规则
	val.AddRule("name", "用户名","required,len","2-20")
	info := val.CheckInfo()
	if info != nil {
		u.JSON(http.StatusOK, info)
		return
	}

	ss := user.Create(values)
	u.JSON(http.StatusOK, ss)
}
