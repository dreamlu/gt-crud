// author: dreamlu
package controllers

import (
	"demo/models"
	"github.com/dreamlu/deercoder-gin/util/xss"
	"github.com/gin-gonic/gin"
	"net/http"
)

var q models.Order
//用户信息分页
func GetOrderBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	xss.XssMap(values)
	ss := q.GetMoreBySearch(values)
	u.JSON(http.StatusOK, ss)
}