// author: dreamlu
package controllers

import (
	"demo/models/order"
	"demo/util/cm"
	"github.com/gin-gonic/gin"
	"net/http"
)

var q order.Order

//用户信息分页
func GetOrderBySearch(u *gin.Context) {
	ss := q.GetMoreBySearch(cm.ToCMap(u))
	u.JSON(http.StatusOK, ss)
}
