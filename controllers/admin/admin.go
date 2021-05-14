package admin

import (
	"demo/models/admin"
	"demo/util/result"
	"demo/util/token"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var p admin.Admin

//根据id获得data
func Get(u *gin.Context) {
	u.JSON(http.StatusOK, result.ResGet(p.Get(result.ToCMap(u))))
}

//data信息分页
func Search(u *gin.Context) {
	u.JSON(http.StatusOK, result.ResPager(p.Search(result.ToCMap(u))))
}

//data信息删除
func Delete(u *gin.Context) {
	id := u.Param("id")
	err := p.Delete(id)
	u.JSON(http.StatusOK, result.Res(err))
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

	err = p.Update(&data)
	u.JSON(http.StatusOK, result.Res(err))
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
	u.JSON(http.StatusOK, result.Res(err))
}

type LoginAdmin struct {
	admin.Admin
	Key  string `json:"key"`
	Code string `json:"code"`
}

// 登录
func Login(u *gin.Context) {

	var (
		login   LoginAdmin
		sqlData admin.Admin
	)

	// 自定义日期格式问题
	err := u.ShouldBindJSON(&login)
	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err))
		return
	}

	// 验证码验证,有需要打开验证
	//if !captcha.Check(login.Key, login.Code) {
	//	u.JSON(http.StatusOK, result.GetText("验证码不正确"))
	//  return
	//}

	// 查找
	err = gt.NewCrud(gt.Data(&sqlData)).Select("select id,password,role from admin where name = ?", login.Name).Single().Error()
	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	// 验证不通过
	if sqlData.Password != util.AesEn(login.Password) {
		u.JSON(http.StatusOK, result.GetText("账号或密码错误"))
		return
	}

	tk := token.GetToken(sqlData.ID, 0)

	u.JSON(http.StatusOK, result.MapSuccess.
		Add("id", sqlData.ID).
		Add("token", tk).
		Add("role", sqlData.Role))
}
