package db

/*made by lucheng*/
import (
	"fmt"
	"kpx_crm/conf"
	"kpx_crm/util"
	"kpx_crm/util/lib"
	"reflect"
	"strconv"
	"strings"
)

/*根据model中表模型的json标签获取表字段,增加select where 条件 关键字查询*/
//尝试单条件搜索全部,多条件搜索关系为且
func GetSearchSqlByKey(model interface{},keys []string) (sql string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer

	for _,v := range keys{
		sql += "("

		for i := 0; i < typ.NumField(); i++ {
			tag := typ.Field(i).Tag.Get("json")
			//buffer.WriteString("`"+tag + "`,")
			sql += "`" + tag + "` like binary '%" + v + "%' or "
		}
		sql = string([]byte(sql)[:len(sql)-3]) //去掉or
		//多条件and关系
		sql += ") and "
	}

	sql = string([]byte(sql)[:len(sql)-4]) //去掉and
	return sql
}



/*根据model中表模型的json标签获取表字段,将select* 变为对应的字段名,增加别名,表连接问题*/
func GetColumnsSql(model interface{},alias string) (sql string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		//buffer.WriteString("`"+tag + "`,")
		sql += alias+".`" + tag + "`,"
	}
	sql = string([]byte(sql)[:len(sql)-1]) //去掉点,
	return sql
}

/*根据model中表模型的json标签获取表字段,将select* 变为对应的字段名*/
func GetSqlColumnsSql(model interface{}) (sql string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		//buffer.WriteString("`"+tag + "`,")
		sql += "`" + tag + "`,"
	}
	sql = string([]byte(sql)[:len(sql)-1]) //去掉点,
	return sql
}

/*传入表名,查询语句拼接*/
func SearchTableSql(model interface{}, tablename string, args map[string][]string) (sqlnolimit, sql string, clientPage, everyPage int) {

	//页码,每页数量
	clientPageStr := conf.GetConfigValue("clientPage") //默认第1页
	everyPageStr := conf.GetConfigValue("everyPage")   //默认10页

	//尝试将select* 变为对应的字段名
	sql = fmt.Sprintf("select %s from `%s` where 1=1 and ", GetSqlColumnsSql(model), tablename)
	for k, v := range args {
		if k == "clientPage" {
			clientPageStr = v[0]
			continue
		}
		if k == "everyPage" {
			everyPageStr = v[0]
			continue
		}
		if v[0] == "" { //条件值为空,舍弃
			continue
		}
		v[0] = strings.Replace(v[0], "'", "\\'", -1) //转义
		sql += k + " like '%" + v[0] + "%' and "
	}

	clientPage, _ = strconv.Atoi(clientPageStr)
	everyPage, _ = strconv.Atoi(everyPageStr)

	c := []byte(sql)
	sql = string(c[:len(c)-4]) //去and
	sqlnolimit = strings.Replace(sql, GetSqlColumnsSql(model), "count(id) as sum_page", -1)
	sql += "order by id desc limit " + strconv.Itoa((clientPage-1)*everyPage) + "," + everyPageStr

	return sqlnolimit, sql, clientPage, everyPage
}

/*传入表名,查询语句拼接,包含一个日期范围,time:被比较的表中的字段名*/
func SearchTableSqlInclueDate(time string, tablename string, args map[string][]string) (sqlnolimimt, sql string, clientPage, everyPage int) {

	//页码,每页数量
	clientPageStr := conf.GetConfigValue("clientPage") //默认第1页
	everyPageStr := conf.GetConfigValue("everyPage")   //默认10页

	sql = "select * from `" + tablename + "` where 1=1 and "
	for k, v := range args {
		if k == "time1" { //时间范围
			sql += "timestampdiff(day,'" + v[0] + "'," + time + ") >= 0 and "
			continue
		}
		if k == "time2" {
			sql += "timestampdiff(day,'" + v[0] + "'," + time + ") <= 0 and "
			continue
		}
		if k == "clientPage" { //页码
			clientPageStr = v[0]
			continue
		}
		if k == "everyPage" { //每页数量
			everyPageStr = v[0]
			continue
		}
		if v[0] == "" { //条件为空,舍弃
			continue
		}
		sql += k + " like '%" + v[0] + "%' and "
	}

	clientPage, _ = strconv.Atoi(clientPageStr)
	everyPage, _ = strconv.Atoi(everyPageStr)

	c := []byte(sql)
	sql = string(c[:len(c)-4]) //去and
	sqlnolimimt = sql

	sql += "order by id desc limit " + strconv.Itoa((clientPage-1)*everyPage) + "," + everyPageStr

	return sqlnolimimt, sql, clientPage, everyPage

}

/*
传入数据库表名
更新语句拼接
*/
func GetUpdateSqlById(tablename string, args map[string][]string) (sql, id string) {

	sql = "update `" + tablename + "` set "
	for k, v := range args {
		if k == "id" {
			id = v[0]
			continue
		}
		//密码aes加密
		if k == "userpassword" {
			v[0] = util.AesEn(v[0])
		}
		v[0] = strings.Replace(v[0], "'", "\\'", -1)
		sql += "`" + k + "`='" + v[0] + "',"
	}
	c := []byte(sql)
	sql = string(c[:len(c)-1]) //去掉点
	sql += " where id=" + id
	return sql, id
}

/*
传入数据库表名
插入语句拼接
*/
func GetInsertSql(tablename string, args map[string][]string) (sql string) {

	//sql拼接
	var values []string
	sql = "insert `" + tablename + "`("
	for k, v := range args {
		//密码aes加密
		if k == "userpassword" {
			v[0] = util.AesEn(v[0])
		}
		sql += "`"+k + "`,"
		values = append(values, v[0])
	}
	c := []byte(sql)
	sql = string(c[:len(c)-1]) + ") value(" //去掉点

	for _, v := range values {
		v = strings.Replace(v, "'", "\\'", -1) //转义
		sql += "'" + v + "',"
	}
	c = []byte(sql)
	sql = string(c[:len(c)-1]) + ")" //去掉点

	return sql
}

/*============================================================================*/
/*==========================增删改查通用=========made=by=lucheng================*/
/*============================================================================*/

//获得数据,根据sql语句
func GetDataBySql(data interface{}, sql string, args...string) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	dba := DB.Raw(sql,args).Scan(data)
	num := dba.RowsAffected

	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = 200
		getinfo.Msg = "请求成功"
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

//获得数据,根据name条件
func GetDataByName(data interface{}, name,value string) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	dba := DB.Where(name + " = ?",value).Find(data) //只要一行数据时使用 LIMIT 1,增加查询性能
	num := dba.RowsAffected

	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = 200
		getinfo.Msg = "请求成功"
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

//获得数据,根据id
func GetDataById(data interface{}, id string) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	dba := DB.First(data, id) //只要一行数据时使用 LIMIT 1,增加查询性能
	num := dba.RowsAffected

	//有数据是返回相应信息
	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = 200
		getinfo.Msg = "请求成功"
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

//获得数据,分页/查询
func GetDataBySearch(model, data interface{}, tablename string, args map[string][]string) interface{} {
	var info interface{}
	//dest = dest.(reflect.TypeOf(dest).Elem())//type != type?
	var getinfo lib.GetInfo

	sqlnolimit, sql, clientPage, everyPage := SearchTableSql(model, tablename, args)

	dba := DB.Raw(sqlnolimit).Scan(&getinfo.Pager)
	num := getinfo.Pager.SumPage
	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//DB.Debug().Find(&dest)
		dba = DB.Raw(sql).Scan(data)
		if dba.Error != nil {
			info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
			return info
		}

		//统计页码等状态
		getinfo.Status = 200
		getinfo.Msg = "请求成功"
		getinfo.Data = data //数据
		//getinfo.Pager.SumPage = num
		getinfo.Pager.ClientPage = clientPage
		getinfo.Pager.EveryPage = everyPage

		info = getinfo
	}
	return info
}

/*删除通用,任意参数*/
func DeleteDataByName(tablename string, key, value string) interface{} {
	var info interface{}
	sql := "delete from `" + tablename + "` where " + key + "=?"

	//sql = string([]byte(sql)[:len(sql)-5])
	dba := DB.Exec(sql, value)

	num := dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapDelete
	}
	return info
}

/*修改数据,通用*/
func UpdateData(tablename string, args map[string][]string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	sql, _ := GetUpdateSqlById(tablename, args)

	dba := DB.Exec(sql)
	num = dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapUpdate
	}
	return info
}

/*创建数据,通用*/
func CreateData(tablename string, args map[string][]string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	sql := GetInsertSql(tablename, args)
	dba := DB.Exec(sql)
	num = dba.RowsAffected

	if dba.Error != nil {
		info = lib.GetMapDataError(lib.CodeSql, dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapError
	} else {
		info = lib.MapCreate
	}
	return info
}
