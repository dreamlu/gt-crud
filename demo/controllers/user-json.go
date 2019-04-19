// author:  dreamlu
package controllers

import (
	"demo/models"
	"github.com/dreamlu/deercoder-gin/util/xss"
	"github.com/gin-gonic/gin"
	"net/http"
)

//根据id获得用户获取
func GetByIdJ(u *gin.Context) {
	id := u.Query("id")
	ss := p.GetByIDJ(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息分页
func GetBySearchJ(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	ss := p.GetBySearchJ(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func DeleteJ(u *gin.Context) {
	id := u.Param("id")
	ss := p.DeleteJ(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func UpdateJ(u *gin.Context) {
	var user models.User

	_ = u.ShouldBindJSON(&user)
	// do something

	ss := p.UpdateJ(&user)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func CreateJ(u *gin.Context) {
	var user models.User

	_ = u.ShouldBindJSON(&user)

	ss := p.CreateJ(&user)
	u.JSON(http.StatusOK, ss)
}
