package client

import (
	"demo/util/models"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// token
func Token(u *gin.Context) {
	client_id := u.Query("id")
	if client_id == "" {
		u.JSON(http.StatusOK, result.GetError("id不能为空"))
		return
	}
	ca := cache.NewCache()
	var model models.TokenModel
	model.ID, _ = strconv.ParseUint(client_id, 10, 64)
	newID, _ := id.NewID(1)
	model.Token = newID.String()
	ca.Set(model.Token, cache.CacheModel{
		Time: cache.CacheDay,
		Data: model,
	})

	u.JSON(http.StatusOK, result.MapSuccess.Add("token", model.Token))
}
