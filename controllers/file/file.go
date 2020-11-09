package file

import (
	File "github.com/dreamlu/gt/tool/file"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"image"
	"net/http"
	"strings"
)

// 单文件上传
// use gin upload file
func UploadFile(u *gin.Context) {

	name := u.PostForm("name") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	upFile := File.File{
		Name: name,
	}
	path, err := upFile.GetUploadFile(file)
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	f, err := file.Open()
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
	}
	defer f.Close()
	filenameSplit := strings.Split(file.Filename, ".")
	fType := filenameSplit[len(filenameSplit)-1]
	fType = strings.ToLower(fType)
	switch fType {
	case "jpeg", "jpg", "png":
		c, _, err := image.DecodeConfig(f)
		if err != nil {
			u.JSON(http.StatusOK, result.CError(err))
			return
		}
		u.JSON(http.StatusOK, result.MapSuccess.Add("path", path).Add("width", c.Width).Add("height", c.Height))
	default:
		u.JSON(http.StatusOK, result.MapSuccess.Add("path", path))
	}
}

type Path struct {
	Path string `json:"path"`
}

// 多文件上传
// use gin upload file
func UploadMultiFile(u *gin.Context) {

	var paths []Path
	form, err := u.MultipartForm()
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}
	files := form.File["file"]

	var path Path
	for _, file := range files {

		upFile := File.File{}
		path.Path, err = upFile.GetUploadFile(file)
		if err != nil {
			u.JSON(http.StatusOK, result.CError(err))
			return
		}
		paths = append(paths, path)
	}

	u.JSON(http.StatusOK, result.GetSuccess(paths))
}
