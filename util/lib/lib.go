package lib

/*made by lucheng*/
/*分页数据信息*/
type GetInfo struct {
	GetInfoN
	Pager  Pager       `json:"pager"`
}

type Pager struct {
	ClientPage int64 `json:"clientpage"` //当前页码
	SumPage    int64 `json:"sumpage"`    //总页数
	EveryPage  int64 `json:"everypage"`  //每一页显示的数量
}

/*无分页数据信息*/
/*分页数据信息*/
type GetInfoN struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"` //数据,通用接口
}

/*全局变量*/
var MapError = map[string]interface{}{"status": 500, "msg": "系统繁忙"}
var MapUpdate = map[string]interface{}{"status": 206, "msg": "修改成功"}
var MapDelete = map[string]interface{}{"status": 209, "msg": "删除成功"}
var MapCreate = map[string]interface{}{"status": 201, "msg": "创建成功"}
var MapNoResult = map[string]interface{}{"status": 204, "msg": "暂无数据"}
var MapNoAuth = map[string]interface{}{"status": 203, "msg": "请求非法"}
var MapNoToken = map[string]interface{}{"status": 213, "msg": "用户凭证失效,请重新登录"}
var MapCaptcha = map[string]interface{}{"status": 214, "msg": "验证码验证失败"}
var MapExistOrNo = map[string]interface{}{"status": 220, "msg": "数据已存在 or 条件不存在"}
var MapCountErr = map[string]interface{}{"status": 211, "msg": "用户账号或密码错误"}
var MapNoCount = map[string]interface{}{"status": 211, "msg": "用户账号不存在"}
var MapNoArgs = map[string]interface{}{"status": 223, "msg": "缺少参数"}

/*缺省或验证字段*/
var MapPhone = map[string]interface{}{"status": 210, "msg": "字段提交不合法", "data": "电话号码格式非法"}
var MapEmail = map[string]interface{}{"status": 210, "msg": "字段提交不合法", "data": "邮箱格式非法"}
var MapEmpty = map[string]interface{}{"status": 210, "msg": "字段提交不合法", "data": "字段内容不能为空"}

/*微信小程序*/
//var WxEcryptError = map[string]interface{}{"status": 230, "msg": "缺少参数"}

/*约定状态码*/
const (
	CodeSuccess    = 200 //请求成功
	CodeRequired   = 210 //必填项
	CodeSql        = 222 //数据库错误统一状态
	CodeFile       = 224 //文件上传相关
	CodeNoDelete   = 225 //存在外健约束(逻辑或数据库约束)
	CodeEcrypt     = 230 //数据解密失败
	CodeWx         = 240 //小程序相关
	CodeWxPay      = 242 //支付失败
	CodeWxWithDraw = 243 //提现失败
	CodeOrder      = 251 //订单相关
	CodeValidator  = 255 //验证相关
	CodeAliPay     = 262 //支付宝支付失败
	CodeError      = 270 //通用错误信息
	CodeText       = 271 //全局文字错误提示
	CodeChat       = 280 //chat相关
)

/*约定提示信息*/
const (
	MsgSuccess = "请求成功"
)

/*错误信息,通用*/
type MapData struct {
	Status int64       `json:"status"`
	Msg    interface{} `json:"msg"`
}

/*错误信息通用,状态码及信息提示*/
func GetMapData(status int64, msg interface{}) MapData{
	me := MapData{
		Status: status,
		Msg:    msg,
	}
	return me
}
