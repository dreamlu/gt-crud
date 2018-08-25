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
