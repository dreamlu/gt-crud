package captcha

import (
	"demo/util/result"
	"github.com/afocus/captcha"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/log"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
)

// key: uuid
// 验证码,5分钟
func Captcha(u *gin.Context) {

	key := u.Query("key")
	if key == "" {
		u.JSON(http.StatusOK, result.GetText("key不能为空"))
		return
	}
	c := captcha.New()
	// 设置字体
	c.SetFont("conf/my.ttf")
	// 创建验证码 4个字符 captcha.NUM 字符模式数字类型
	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
	img, code := c.Create(4, captcha.ALL)
	ce := cache.NewCache()
	err := ce.Set(key, cache.CacheModel{
		Time: 5 * cache.CacheMinute,
		Data: code,
	})
	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err))
		return
	}

	png.Encode(u.Writer, img)
}

// code: 需要验证的code
func Check(key, code string) bool {
	if key == "" {
		return false
	}
	ce := cache.NewCache()
	cd, err := ce.Get(key)
	if err != nil {
		log.Error(err)
		return false
	}
	if cd.Data.(string) == code {
		return true
	}
	return false
}
