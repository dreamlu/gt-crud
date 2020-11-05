package client

import (
	"demo/controllers"
	"demo/models"
	"demo/models/client"
	"demo/util/models/token"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 自定义额外接口
func New(model interface{}, arrayModel interface{}) controllers.ComController {
	return controllers.ComController{Service: &models.Com{
		Model:         model,
		ArrayModel:    arrayModel,
		UpdateService: &client.Client{},
	}}
}

// token
func Token(u *gin.Context) {
	client_id := u.Query("id")
	if client_id == "" {
		u.JSON(http.StatusOK, result.GetError("id不能为空"))
		return
	}
	ca := cache.NewCache()
	var model token.TokenModel
	model.ID, _ = strconv.ParseUint(client_id, 10, 64)
	newID, _ := id.NewID(1)
	model.Token = newID.String()
	ca.Set(model.Token, cache.CacheModel{
		Time: cache.CacheDay,
		Data: model,
	})

	u.JSON(http.StatusOK, result.MapSuccess.Add("token", model.Token))
}
