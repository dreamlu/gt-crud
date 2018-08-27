package lib

/*made by lucheng*/
/*分页数据信息*/
type GetInfo struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"` //数据,通用接口
	Pager  Pager       `json:"pager"`
}

type Pager struct {
	ClientPage int `json:"clientpage"` //当前页码
	SumPage    int `json:"sumpage"`    //总页数
	EveryPage  int `json:"everypage"`  //每一页显示的数量
}

/*无分页数据信息*/
/*分页数据信息*/
type GetInfoN struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"` //数据,通用接口
}

/*全局变量*/
var MapError = map[string]string{"status": "500", "msg": "系统繁忙"}
var MapUpdate = map[string]string{"status": "206", "msg": "修改成功"}
var MapDelete = map[string]string{"status": "209", "msg": "删除成功"}
var MapCreate = map[string]string{"status": "201", "msg": "创建成功"}
var MapNoResult = map[string]string{"status": "204", "msg": "暂无数据"}
var MapNoAuth = map[string]string{"status": "203", "msg": "请求非法"}
var MapExistOrNo = map[string]string{"status": "220", "msg": "数据已存在 or 条件值不存在"}

/*缺省或验证字段*/
var MapPhone = map[string]string{"status": "210", "msg": "手机号码格式非法"}
var MapEmail = map[string]string{"status": "210", "msg": "邮箱格式非法"}
var MapEmpty = map[string]string{"status": "210", "msg": "字段内容不能为空"}

/*约定状态码*/
const (
	CodeSql      = 222
	CodeRequired = 210
)

/*错误信息,通用*/
type MapDataError struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

/*错误信息通用,状态码及信息提示*/
func GetMapDataError(status int, msg interface{}) MapDataError {
	me := MapDataError{
		Status: status,
		Msg:    msg,
	}
	return me
}
