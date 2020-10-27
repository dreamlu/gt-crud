package applet

import (
	"demo/models/admin/applet"
	"demo/util/file"
	"encoding/json"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

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
