package models

import (
	"deercoder-gin/util"
	"deercoder-gin/util/db"
	"deercoder-gin/util/lib"
)

/*用户信息模型*/
type Userinfo struct {
	ID         uint          `json:"id"`
	Userinfo   string        `json:"userinfo"`
	Updatetime util.JsonTime `json:"updatetime"`
	UserID     uint          `json:"user_id"`
}

/*用户信息输出模型,属于关系*/
type UserinfoOut struct {
	User
	Userinfo []*Userinfo `json:"userinfo"`
}

//获得用户信息以及相关数据,分页
func GetUserinfoBySearch(args map[string][]string) interface{} {
	var info interface{}
	//dest = dest.(reflect.TypeOf(dest).Elem())//type != type?
	var getinfo lib.GetInfo

	sqlnolimit, sql, clientPage, everyPage := db.SearchTableSql(User{}, "user", args)

	dba := db.DB.Raw(sqlnolimit).Scan(&getinfo.Pager)
	num := getinfo.Pager.SumPage
	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//DB.Debug().Find(&dest)
		var user []*UserinfoOut
		dba = db.DB.Raw(sql).Scan(&user)
		if dba.Error != nil {
			info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
			return info
		}
		//查询下面对应数据
		for _, v := range user {
			db.DB.Where("user_id = ?",v.ID).Find(&v.Userinfo)
		}

		//统计页码等状态
		getinfo.Status = 200
		getinfo.Msg = "请求成功"
		getinfo.Data = user //数据
		getinfo.Pager.ClientPage = clientPage
		getinfo.Pager.EveryPage = everyPage

		info = getinfo
	}
	return info
}
