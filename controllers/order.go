// author: dreamlu
package controllers

import (
	"demo/models"
	"demo/util/cm"
	"github.com/gin-gonic/gin"
	"net/http"
)

var q models.Order

//用户信息分页
func GetOrderBySearch(u *gin.Context) {
	ss := q.GetMoreBySearch(cm.ToCMap(u))
	u.JSON(http.StatusOK, ss)
}
