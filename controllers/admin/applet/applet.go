package applet

import (
	"demo/models/admin/applet"
	"demo/util/cm"
	"demo/util/file"
	"encoding/json"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var p applet.Applet

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
		data applet.Applet
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
		data applet.Applet
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

// 下载代码处理
type ProjectConfig struct {
	Description               string      `json:"description"`
	PackOptions               interface{} `json:"pack_options"`
	Setting                   interface{} `json:"setting"`
	CompileType               interface{} `json:"compile_type"`
	LibVersion                interface{} `json:"lib_version"`
	Appid                     interface{} `json:"appid"`
	Projectname               interface{} `json:"projectname"`
	CloudfunctionTemplateRoot interface{} `json:"cloudfunction_template_root"`
	WatchOptions              interface{} `json:"watch_options"`
	DebugOptions              interface{} `json:"debug_options"`
	Scripts                   interface{} `json:"scripts"`
	SimulatorType             interface{} `json:"simulator_type"`
	SimulatorPluginLibVersion interface{} `json:"simulator_plugin_lib_version"`
	Condition                 interface{} `json:"condition"`
}

// 下载代码
func DownLoad(u *gin.Context) {
	var (
		wx applet.Applet
	)

	_ = u.ShouldBindJSON(&wx)
	if err := wx.GetByAdminID(wx.AdminID); err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	err := rwConfig(wx.Appid)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	u.JSON(http.StatusOK, result.MapSuccess.Add("path", gt.Configger().GetString("app.staticpath")+"app/dist.zip"))
}

func rwConfig(appid string) error {
	// 读取文件
	dir := gt.Configger().GetString("app.staticpath") + "app/dist/"
	path := dir + "project.config.json"
	by, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var data ProjectConfig
	err = json.Unmarshal(by, &data)
	if err != nil {
		return err
	}
	data.Appid = appid
	by, err = json.Marshal(data)
	if err != nil {
		return err
	}
	// 写入文件
	err = ioutil.WriteFile(path, by, os.ModePerm)
	if err != nil {
		return err
	}

	// 压缩文件
	newZip := gt.Configger().GetString("app.staticpath") + "app/dist.zip"
	err = file.Zip(dir, newZip)
	if err != nil {
		return err
	}

	return nil
}
