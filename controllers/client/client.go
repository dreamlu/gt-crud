package client

import (
	"demo/models/client"
	"demo/util/cm"
	"demo/util/models"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var p client.Client

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

//根据id获得data
func Get(u *gin.Context) {
	data, err := p.Get(cm.ToCMap(u))
	u.JSON(http.StatusOK, cm.ResGet(err, data))
}

//data信息分页
func Search(u *gin.Context) {
	datas, pager, err := p.Search(cm.ToCMap(u))
	u.JSON(http.StatusOK, cm.ResPager(err, datas, pager))
}

//data信息删除
func Delete(u *gin.Context) {
	id := u.Param("id")
	err := p.Delete(id)
	u.JSON(http.StatusOK, cm.Res(err))
}

//data信息修改
func Update(u *gin.Context) {
	var (
		data client.Client
	)
	// json 类型需要匹配
	// 与spring boot不同
	// 不能自动将字符串转成对应类型
	// 严格匹配
	err := u.ShouldBindJSON(&data)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	// do something

	err = p.Update(&data)
	u.JSON(http.StatusOK, cm.Res(err))
}

//新增data信息
func Create(u *gin.Context) {
	var (
		data client.Client
	)
	// 自定义日期格式问题
	err := u.ShouldBindJSON(&data)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	_, err = p.Create(&data)
	u.JSON(http.StatusOK, cm.Res(err))
}
