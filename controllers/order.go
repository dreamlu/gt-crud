// author: dreamlu
package controllers

import (
	"demo/models"
	"demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var q models.Order

//用户信息分页
func GetOrderBySearch(u *gin.Context) {
	ss := q.GetMoreBySearch(util.ToCMap(u))
	u.JSON(http.StatusOK, ss)
}
