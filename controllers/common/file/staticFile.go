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

// static file path
func StaticFile(u *gin.Context) {

	var fileByte []byte

	// get the params
	url := u.Request.URL.Path
	urlSp := strings.Split(url, "/")
	name := urlSp[len(urlSp)-1]
	width, err := strconv.Atoi(u.Query("width"))
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(u.Query("height"))
	if err != nil {
		width = 0
	}

	// upload dir
	uploadDir := conf.GetString("app.filepath")

	// 文件查找 是否存在 不存在则压缩
	if width != 0 || height != 0 {
		newFileName := strconv.Itoa(width) + "-" + strconv.Itoa(height) + "-" + name
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
		}
	} else {
		// 文件读取
		fileByte, err = ioutil.ReadFile(uploadDir + name)
		if err != nil {
			u.JSON(http.StatusOK, result.GetError(err))
			return
		}
	}

	// return file
	u.Data(http.StatusOK, "utf-8", fileByte)
}
