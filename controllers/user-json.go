// @author  dreamlu
package controllers

import (
	"demo/models"
	"github.com/dreamlu/go-tool/util/xss"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//根据id获得用户获取
func GetByIdJ(u *gin.Context) {
	id := u.Query("id")
	ss := p.GetByIDJ(id)
	u.JSON(http.StatusOK, ss)
}

// 用户信息分页
func GetBySearchJ(u *gin.Context) {
	// this is get url 参数
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

	// json 类型需要匹配
	// 与spring boot不同
	// 不能自动将字符串转成对应类型
	// 严格匹配
	_ = u.ShouldBindJSON(&user)
	// do something

	ss := p.UpdateJ(&user)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func CreateJ(u *gin.Context) {
	var user  models.User

	// 自定义日期格式问题
	err := u.ShouldBindJSON(&user)
	log.Println(err)


	ss := p.CreateJ(&user)
	u.JSON(http.StatusOK, ss)
}
