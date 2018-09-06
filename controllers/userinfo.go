package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"deercoder-gin/models"
)

//根据条件获得用户,分页
func GetUserinfoBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form
	ss := models.GetUserinfoBySearch(values)
	u.JSON(http.StatusOK, ss)
}
