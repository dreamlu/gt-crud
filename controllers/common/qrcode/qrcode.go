package qrcode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 普通二维码生成
func PQrcode(u *gin.Context) {
	url := u.Query("url")
	w := u.Writer

	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}()
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(png)))
	w.Write(png)
}
