package file

import (
	File "github.com/dreamlu/gt/tool/file"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 单文件上传
// use gin upload file
func UploadFile(u *gin.Context) {

	name := u.PostForm("name") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
	}
	upFile := File.File{
		Name: name,
	}
	path, err := upFile.GetUploadFile(file)
	u.JSON(http.StatusOK, map[string]interface{}{result.Status: result.CodeSuccess, result.Msg: result.MsgSuccess, "path": path})
}
