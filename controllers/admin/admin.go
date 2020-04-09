package admin

import (
	"demo/models/admin"
	"demo/util/cm"
	"demo/util/models"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var p admin.Admin

//根据id获得data
func Get(u *gin.Context) {
	id := u.Query("id")
	data, err := p.Get(id)
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
		data admin.Admin
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

	_, err = p.Update(&data)
	u.JSON(http.StatusOK, cm.Res(err))
}

//新增data信息
func Create(u *gin.Context) {
	var (
		data admin.Admin
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

// 登录
func Login(u *gin.Context) {
	var login, sqlData admin.Admin

	// 自定义日期格式问题
	err := u.ShouldBindJSON(&login)
	log.Println(err)

	// 查找
	gt.NewCrud(gt.Data(&sqlData)).Select("select id,password,role from admin where name = ?", login.Name).Single()

	// 验证不通过
	if sqlData.Password != util.AesEn(login.Password) {
		u.JSON(http.StatusOK, result.MapCountErr)
		return
	}

	ca := cache.NewCache()
	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	var model models.TokenModel
	model.ID = sqlData.ID
	newID, _ := id.NewID(1)
	model.Token = newID.String()
	ca.Set(model.Token, cache.CacheModel{
		Time: cache.CacheDay,
		Data: model,
	})

	u.JSON(http.StatusOK, result.MapSuccess.
		Add("id", model.ID).
		Add("token", model.Token).
		Add("role", sqlData.Role))
}
