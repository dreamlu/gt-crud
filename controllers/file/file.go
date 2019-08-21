package file

import (
	File "github.com/dreamlu/go-tool/tool/file"
	"github.com/dreamlu/go-tool/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 单文件上传
// use gin upload file
func UploadFile(u *gin.Context) {

	name := u.PostForm("name") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		u.JSON(http.StatusOK, result.MapData{Status: result.CodeFile, Msg: err.Error()})
	}
	upFile := File.File{
		Name: name,
	}
	path, err := upFile.GetUploadFile(file)
	u.JSON(http.StatusOK, map[string]interface{}{result.Status: result.CodeFile, result.Msg: result.MsgFile, "path": path})
}
