package file

import (
	"demo/util/result"
	"github.com/dreamlu/gt/tool/conf"
	"github.com/dreamlu/gt/tool/file"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// StaticFile static file path
func StaticFile(u *gin.Context) {
	width, err := strconv.Atoi(u.Query("width"))
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(u.Query("height"))
	if err != nil {
		width = 0
	}

	if width == 0 && height == 0 {
		return
	}

	var (
		fileByte     []byte
		staticDir    = conf.GetString("app.filepath")
		url          = u.Request.URL.Path
		urlSp        = strings.Split(url, staticDir)
		dirFileName  = urlSp[len(urlSp)-1]
		dirFileNameS = strings.Split(dirFileName, "/")
		name         = dirFileNameS[len(dirFileNameS)-1]
		dayDir       = ""
		newFileName  = strconv.Itoa(width) + "-" + strconv.Itoa(height) + "-" + name
	)

	// upload dir
	if len(dirFileNameS) > 1 {
		dayDir = dirFileNameS[0]
	}
	uploadDir := staticDir + dayDir + "/"

	// 文件查找 是否存在 不存在则压缩
	fileByte, err = ioutil.ReadFile(uploadDir + newFileName)
	if err != nil { // 文件不存在
		// 文件压缩
		fileUtil := &file.File{
			Path:    uploadDir + name,
			NewPath: uploadDir + newFileName,
			Width:   width,
			Height:  height,
		}
		err = fileUtil.CompressImage("jpg")
		fileByte, err = ioutil.ReadFile(fileUtil.NewPath)
		if err != nil {
			u.JSON(http.StatusOK, result.GetError(err))
			return
		}
		u.Data(http.StatusOK, "utf-8", fileByte)
	}
}
