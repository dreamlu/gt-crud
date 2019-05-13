// @author  dreamlu
package lib

// status and msg
const (
	Status = "status"
	Msg    = "msg"
)

// 约定状态码
// 或 通过GetMapData()自定义
const (
	CodeSuccess    = 200 // 请求成功
	CodeCreate     = 201 // 创建成功
	CodeNoAuth     = 203 // 请求非法
	CodeNoResult   = 204 // 暂无数据
	CodeUpdate     = 206 // 修改成功
	CodeDelete     = 209 // 删除成功
	CodeValidator  = 210 // 字段验证
	CodeCount      = 211 // 账号相关
	CodeCaptcha    = 214 // 验证码
	CodeValSuccess = 217 // 验证成功
	CodeValError   = 218 // 验证失败
	CodeExistOrNo  = 220 // 数据无变化
	CodeSQL        = 222 // 数据库相关
	CodeLackArgs   = 223 // 缺少参数
	CodeFile       = 224 // 文件上传相关
	CodeNoDelete   = 225 // 存在外健约束(逻辑或数据库约束)
	CodeEcrypt     = 230 // 数据解密失败
	CodeWx         = 240 // 微信小程序相关
	CodeWxPay      = 242 // 微信支付相关
	CodeWxWithDraw = 243 // 微信提现相关
	CodeOrder      = 251 // 订单相关
	CodeAliPay     = 262 // 支付宝支付相关
	CodeText       = 271 // 全局文字提示
	CodeChat       = 280 // chat相关
	CodeError      = 500 // 系统繁忙
)

// 约定提示信息
const (
	MsgSuccess    = "请求成功"
	MsgCreate     = "创建成功"
	MsgNoAuth     = "请求非法"
	MsgNoResult   = "暂无数据"
	MsgDelete     = "删除成功"
	MsgUpdate     = "修改成功"
	MsgError      = "未知错误"
	MsgCaptcha    = "验证码验证失败"
	MsgExistOrNo  = "数据无变化"
	MsgCountErr   = "用户账号或密码错误"
	MsgNoCount    = "用户账号不存在"
	MsgLackArgs   = "缺少参数"
	MsgValSuccess = "验证成功"
	MsgValError   = "验证失败"
	MsgFile       = "上传成功"
)

// 约定提示信息
var (
	MapSuccess    = GetMapData(CodeSuccess, MsgSuccess)       // 请求成功
	MapError      = GetMapData(CodeError, MsgError)           // 通用失败
	MapUpdate     = GetMapData(CodeUpdate, MsgUpdate)         // 修改成功
	MapDelete     = GetMapData(CodeDelete, MsgDelete)         // 删除成功
	MapCreate     = GetMapData(CodeCreate, MsgCreate)         // 创建成功
	MapNoResult   = GetMapData(CodeNoResult, MsgNoResult)     // 暂无数据
	MapNoAuth     = GetMapData(CodeNoAuth, MsgNoAuth)         // 请求非法
	MapCaptcha    = GetMapData(CodeCaptcha, MsgCaptcha)       // 验证码验证失败
	MapExistOrNo  = GetMapData(CodeExistOrNo, MsgExistOrNo)   // 指数据修改没有变化 或者 给的条件值不存在
	MapCountErr   = GetMapData(CodeCount, MsgCountErr)        // 用户账号密码错误
	MapNoCount    = GetMapData(CodeCount, MsgNoCount)         // 用户账号不存在
	MapLackArgs   = GetMapData(CodeLackArgs, MsgLackArgs)     // 缺少参数
	MapValSuccess = GetMapData(CodeValSuccess, MsgValSuccess) // 验证成功
	MapValError   = GetMapData(CodeValError, MsgValError)     // 验证失败
)

// 分页数据信息
type GetInfoPager struct {
	GetInfo
	Pager Pager `json:"pager"`
}

// pager info
type Pager struct {
	ClientPage int64 `json:"client_page"` //当前页码
	SumPage    int64 `json:"sum_page"`    //数据总数量
	EveryPage  int64 `json:"every_page"`  //每一页显示的数量
}

// 无分页数据信息
// 分页数据信息
type GetInfo struct {
	MapData
	Data interface{} `json:"data"` //数据,通用接口
}

// 信息,通用
type MapData struct {
	Status int64       `json:"status"`
	Msg    interface{} `json:"msg"`
}

// 信息通用,状态码及信息提示
func GetMapData(status int64, msg interface{}) MapData {
	me := MapData{
		Status: status,
		Msg:    msg,
	}
	return me
}

// 信息成功通用(成功通用, 无分页)
func GetMapDataSuccess(data interface{}) GetInfo {
	me := GetInfo{
		MapData: MapSuccess,
		Data:    data,
	}
	return me
}

// 信息分页通用(成功通用, 分页)
func GetMapDataPager(data interface{}, clientPage, everyPage, sumPage int64) GetInfoPager {
	me := GetInfoPager{
		GetInfo: GetMapDataSuccess(data),
		Pager: Pager{
			ClientPage: clientPage,
			EveryPage:  everyPage,
			SumPage:    sumPage,
		},
	}
	return me
}

// 信息失败通用
func GetMapDataError(data interface{}) GetInfo {
	me := GetInfo{
		MapData: MapError,
		Data:    data,
	}
	return me
}

// 无分页通用
func GetInfoData(data interface{}, mapData MapData) GetInfo {
	me := GetInfo{
		MapData: mapData,
		Data:    data,
	}
	return me
}

// 分页通用
// 无分页数据
// 统一返回GetInfoPager 分页格式用
func GetInfoPagerData(data interface{}, mapData MapData, pager Pager) GetInfoPager {
	me := GetInfoPager{
		GetInfo: GetInfoData(data, mapData),
		Pager:   pager,
	}
	return me
}

// 分页通用
// 无分页数据
// 统一返回GetInfoPager 分页格式用
func GetInfoPagerDataNil(data interface{}, mapData MapData) GetInfoPager {
	me := GetInfoPager{
		GetInfo: GetInfoData(data, mapData),
	}
	return me
}
