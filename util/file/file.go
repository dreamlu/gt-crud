package file

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//func main() {
//	// 源档案（准备压缩的文件或目录）
//	var src = "log"
//	// 目标文件，压缩后的文件
//	var dst = "log.zip"
//
//	if err := Zip(src, dst); err != nil {
//		log.Fatalln(err)
//	}
//}

func Zip(src, dst string) (err error) {

	// 默认压缩文件内部名(路径)
	zipPaths := strings.Split(src, "/")
	zipPath := zipPaths[len(zipPaths)-2]

	// 创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 下面来将文件写入 zw ，因为有可能会有很多个目录及文件，所以递归处理
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		// 替换文件信息中的文件名
		// 压缩文件中的文件名
		newZipPath := zipPath + strings.Join(strings.Split(path, zipPath)[1:], zipPath)
		//log.Println(newZipPath)
		fh.Name = strings.TrimPrefix(newZipPath, string(filepath.Separator))

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		// 将打开的文件 Copy 到 w
		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}
		// 输出压缩的内容
		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据到目录: %s中\n", path, n, fh.Name)

		return nil
	})
}
