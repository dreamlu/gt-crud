// author:  dreamlu
package controllers

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//根据条件获得用户,分页
func GetUserinfoBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form
	ss := models.GetUserinfoBySearch(values)
	u.JSON(http.StatusOK, ss)
}
