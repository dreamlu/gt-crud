// author:  dreamlu
package file

import (
	"github.com/dreamlu/deercoder-gin"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"github.com/dreamlu/resize"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
	"time"
)

//获得文件上传路径,内部专用
func GetUpoadFile(u *gin.Context) (filename string) {

	fname := u.PostForm("fname") //指定文件名
	file, err := u.FormFile("file")
	if err != nil {
		u.JSON(http.StatusOK, lib.MapData{lib.CodeFile, err.Error()})
	}

	filenameSplit := strings.Split(file.Filename, ".")
	ftype := filenameSplit[len(filenameSplit)-1]
	//防止文件名中多个“.”,获得文件后缀
	filename = "." + ftype
	switch fname {
	case "": //重命名
		filename = time.Now().Format("20060102150405") + filename //时间戳"2006-01-02 15:04:05"是参考格式,具体数字可变(经测试)
	default: //指定文件名
		//防止文件名中多个“.”,获得文件后缀
		filename = fname + filename
	}
	path := deercoder.GetConfigValue("filepath") + filename //文件目录
	u.SaveUploadedFile(file, path)
	switch ftype {
	case "jpeg","jpg","png":
		CompressImage(ftype, path)
	default:
		//处理其他类型文件
	}

	return path
}

//单文件上传
func UpoadFile(u *gin.Context) {

	path := GetUpoadFile(u)
	u.JSON(http.StatusOK, map[string]interface{}{"status": 201, "msg": "创建成功", "filename": path})
}

//图片压缩
func CompressImage(imagetype, path string) error {
	//图片压缩
	var img image.Image
	ImgFile, err := os.Open(path)
	defer ImgFile.Close()
	if err != nil {
		return err
	}
	switch imagetype {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(ImgFile)
		if err != nil {
			return err
		}
	case "png":
		img, err = png.Decode(ImgFile)
		if err != nil {
			return err
		}
	}

	m := resize.Resize(0, 0, img, resize.Lanczos3)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	switch imagetype {
	case "jpeg", "jpg":
		// write new image to file
		jpeg.Encode(out, m, nil)
	case "png":
		png.Encode(out, m) // write new image to file
	}

	return nil
}
