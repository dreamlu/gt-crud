package basic

import (
	"deercoder-gin/conf"
	"deercoder-gin/util/db"
	"deercoder-gin/util/lib"
	"deercoder-gin/util/str"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type Basic struct {
	Address    string `json:"address"`    //ip或域名
	Port       string `json:"port"`       //端口号
	Os         string `json:"os"`         //操作系统
	Goversion  string `json:"goversion"`  //go 版本
	Ginversion string `json:"ginversion"` //gin 版本
	Mysql      string `json:"mysql"`      //mysql版本
	Maxmerory  int64  `json:"maxmerory"`  //最大上传文件大小MB
}

func GetBasicInfo(u *gin.Context) {
	var getinfo lib.GetInfoN
	var basic Basic
	basic.Address = conf.GetConfigValue("domain")
	basic.Port = conf.GetConfigValue("http_port")
	basic.Os = runtime.GOOS
	basic.Goversion = runtime.Version()
	basic.Ginversion = gin.Version
	// router := routers.SetRouter()
	basic.Maxmerory = str.MaxUploadMemory / 1024 / 1024
	db.DB.Raw("select version() as mysql").Scan(&basic)

	getinfo.Data = basic
	getinfo.Status = 200
	getinfo.Msg = "请求成功"
	u.JSON(http.StatusOK, getinfo)
}
