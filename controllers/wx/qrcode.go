package wx

import (
	"demo/models"
	"demo/models/admin/applet"
	"demo/util/cm"
	"fmt"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/result"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/code"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
)

// TODO 约定个时间久的删除key-value, 定时任务
/// ========= 二维码 ===============
type QrCode struct {
	models.IDCom
	Key   string `json:"key" gorm:"varchar(20)"` // key
	Value string `json:"value"`                  // value
}

// 创建二维码key==value转换值
func CreateQrCode(qc *QrCode) {
	//var qc QrCode
	//_ = u.ShouldBindJSON(&qc)
	if qc.Value != "" {
		ids, _ := id.NewID(1)
		qc.Key = ids.String()
	}
	cd := gt.NewCrud(
		gt.Model(QrCode{}),
		gt.Data(&qc),
	).Create()
	if cd.Error() != nil {
		return
	}
}

// 查找
func (c *QrCode) GetByValue() (*QrCode, error) {
	cd := gt.NewCrud(
		gt.Data(&c),
	).Select("select `key` from qr_code where value = ?", c.Value).Single()

	if cd.Error() != nil {
		return nil, cd.Error()
	}

	return c, nil
}

func GetByKey(u *gin.Context) {
	var (
		qc QrCode
	)
	cd := gt.NewCrud(
		gt.Model(QrCode{}),
		gt.Data(&qc),
	).GetByData(cm.ToCMap(u))
	if cd.Error() != nil {
		u.JSON(http.StatusOK, result.CError(cd.Error()))
		return
	}
	u.JSON(http.StatusOK, result.GetSuccess(qc.Value))
}

//二维码,业务量多的情况
func GetQRCode(u *gin.Context) {
	var wx applet.Applet
	_ = u.Request.ParseForm()
	params := u.Request.Form
	if err := wx.GetByAppid(params["appid"][0]); err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	// scene参数转换
	qc := &QrCode{}
	qc.Value = params["value"][0]
	_, err := qc.GetByValue()
	if err != nil {
		if err.Error() == result.MsgNoResult {
			CreateQrCode(qc)
			goto into
		}
		u.JSON(http.StatusOK, result.CError(err))
		return
	}

into:
	coder := code.QRCoder{
		Scene:     qc.Key,            // 参数数据
		Page:      params["page"][0], // 识别二维码后进入小程序的页面链接
		Width:     430,               // 图片宽度
		IsHyaline: false,             // 是否需要透明底色
		AutoColor: true,              // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
		LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
			R: "50",
			G: "50",
			B: "50",
		},
	}

	at := AsToken(wx.Appid, wx.Secret)
	// token: 微信 access_token
	resu, err := coder.UnlimitedAppCode(at.AccessToken)
	defer resu.Body.Close()
	if err != nil {
		u.JSON(http.StatusOK, result.CError(err))
		return
	}

	bodyu, _ := ioutil.ReadAll(resu.Body)
	u.Writer.Header().Add("Content-Type", "image/png")
	u.Writer.Write(bodyu)
}

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
