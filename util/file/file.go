package file

import (
	"deercoder-gin/conf"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"net/http"
	"deercoder-gin/util/lib"
)

//新增用户信息
func UpoadFile(u *gin.Context) {

	file, err := u.FormFile("file")
	if err != nil{
		u.JSON(http.StatusOK,lib.MapError)
	}

	filenameSplit := strings.Split(file.Filename, ".")
	//防止文件名中多个“.”,获得文件后缀
	filename := "." + filenameSplit[len(filenameSplit)-1]
	filename = time.Now().Format("20060102150405") + filename //时间戳"2006-01-02 15:04:05"是参考格式,具体数字可变(经测试)
	path := conf.GetConfigValue("filepath") + filename        //文件目录
	u.SaveUploadedFile(file,path)
	u.JSON(http.StatusOK,map[string]string{"status": "201", "msg": "创建成功", "filename": path})
}
