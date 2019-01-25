package deercoder

/*made by lucheng*/
import (
	"fmt"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"reflect"
	"strconv"
	"strings"
)

// select *替换
// 两张表
func GetDoubleTableColumnsql(model interface{}, table1, table2 string) (sql string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		//table2的数据处理,去除table2_id
		if strings.Contains(tag, table2+"_") && !strings.Contains(tag, table2+"_id") {
			sql += table2 + ".`" + string([]byte(tag)[len(table2)+1:]) + "` as " + tag + "," //string([]byte(tag)[len(table2+1-1):])
			continue
		}
		sql += table1 + ".`" + tag + "`,"
	}
	sql = string([]byte(sql)[:len(sql)-1]) //去掉点,
	return sql
}

// 根据model中表模型的json标签获取表字段
// 将select* 中'*'变为对应的字段名
// 增加别名,表连接问题
func GetColAliasSql(model interface{}, alias string) (sql string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		//buffer.WriteString("`"+tag + "`,")
		sql += alias + ".`" + tag + "`,"
	}
	sql = string([]byte(sql)[:len(sql)-1]) //去掉点,
	return sql
}

// 根据model中表模型的json标签获取表字段
// 将select* 变为对应的字段名
func GetColSql(model interface{}) (sql string) {
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

//=======================================语句拼接==========================================
//========================================================================================

// 两张表名,查询语句拼接
// 表1中有表2 id
func GetDoubleSearchSql(model interface{}, table1, table2 string, args map[string][]string) (sqlnolimit, sql string, clientPage, everyPage int64) {

	//页码,每页数量
	clientPageStr := ClientPageStr
	everyPageStr := EveryPageStr
	var every string
	// 关键字key搜索
	var key string

	//select* 变为对应的字段名
	sql = fmt.Sprintf("select %s from `%s` inner join `%s` on `%s`.%s_id=%s.id where 1=1 and ", GetDoubleTableColumnsql(model, table1, table2), table1, table2, table1, table2, table2)

	sqlnolimit = fmt.Sprintf("select count(%s.id) as sum_page from `%s` inner join `%s` on `%s`.%s_id=%s.id where 1=1 and ", table1, table1, table2, table1, table2, table2)
	for k, v := range args {
		switch k {
		case "clientPage":
			clientPageStr = v[0]
			continue
		case "everyPage":
			everyPageStr = v[0]
			continue
		case "every":
			every = v[0]
			continue
		case "key":
			key = v[0]
			// 多表搜索
			sql, sqlnolimit = lib.GetMoreKeySql(sql, sqlnolimit, key, model , table1+":"+table1, table2+":"+table2)
			//sql, sqlnolimit = lib.GetKeySql(sql, sqlnolimit, key, model , table2)
			continue
		case "":
			continue
		}

		//表2值查询
		if strings.Contains(k, table2+"_") && !strings.Contains(k, table2+"_id") {
			sql += table2 + ".`" + string([]byte(k)[len(table2)+1:]) + "`" + " = '" + v[0] + "' and " //string([]byte(tag)[len(table2+1-1):])
			sqlnolimit += table2 + ".`" + string([]byte(k)[len(table2)+1:]) + "`" + " = '" + v[0] + "' and "
			continue
		}

		v[0] = strings.Replace(v[0], "'", "\\'", -1) //转义
		sql += table1 + "." + k + " = '" + v[0] + "' and "
		sqlnolimit += table1 + "." + k + " = '" + v[0] + "' and "
	}

	clientPage, _ = strconv.ParseInt(clientPageStr, 10, 64)
	everyPage, _ = strconv.ParseInt(everyPageStr, 10, 64)

	sql = string([]byte(sql)[:len(sql)-4])                      //去and
	sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4]) //去and
	if every == "" {
		sql += "order by " + table1 + ".id desc limit " + strconv.FormatInt((clientPage-1)*everyPage, 10) + "," + everyPageStr
	}

	return sqlnolimit, sql, clientPage, everyPage
}

// 传入表名,查询语句拼接
// 单张表
func GetSearchSql(model interface{}, tablename string, args map[string][]string) (sqlnolimit, sql string, clientPage, everyPage int64) {

	//页码,每页数量
	clientPageStr := ClientPageStr
	everyPageStr := EveryPageStr
	every := ""
	// 关键字key搜索
	key := ""

	//将select* 变为对应的字段名
	sql = fmt.Sprintf("select %s from `%s` where 1=1 and ", GetColSql(model), tablename)
	sqlnolimit = fmt.Sprintf("select count(id) as sum_page from `%s` where 1=1 and ", tablename)
	for k, v := range args {
		switch k {
		case "clientPage":
			clientPageStr = v[0]
			continue
		case "everyPage":
			everyPageStr = v[0]
			continue
		case "every":
			every = v[0]
			continue
		case "key":
			key = v[0]
			sql, sqlnolimit = lib.GetKeySql(sql, sqlnolimit, key, model, tablename)
			continue
		case "":
			continue
		}
		v[0] = strings.Replace(v[0], "'", "\\'", -1) //转义
		sql += k + " = '" + v[0] + "' and "          //change 'like' to '='
		sqlnolimit += k + " = '" + v[0] + "' and "
	}

	clientPage, _ = strconv.ParseInt(clientPageStr, 10, 64)
	everyPage, _ = strconv.ParseInt(everyPageStr, 10, 64)

	sql = string([]byte(sql)[:len(sql)-4]) //去and
	sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4]) //去and
	if every == "" {
		sql += "order by id desc limit " + strconv.FormatInt((clientPage-1)*everyPage, 10) + "," + everyPageStr
	}

	return sqlnolimit, sql, clientPage, everyPage
}

// 传入数据库表名
// 更新语句拼接
func GetUpdateSql(tablename string, args map[string][]string) (sql, id string) {

	sql = "update `" + tablename + "` set "
	for k, v := range args {
		if k == "id" {
			id = v[0]
			continue
		}
		v[0] = strings.Replace(v[0], "'", "\\'", -1)//转义
		sql += "`" + k + "`='" + v[0] + "',"
	}
	//c := []byte(sql)
	sql = string([]byte(sql)[:len(sql)-1]) + " where id=" + id //去掉点
	//sql += " where id=" + id
	return sql, id
}

// 传入数据库表名
// 插入语句拼接
func GetInsertSql(tablename string, args map[string][]string) (sql string) {

	//sql拼接
	var values []string
	sql = "insert `" + tablename + "`("
	for k, v := range args {
		sql += "`" + k + "`,"
		values = append(values, v[0])
	}
	sql = string([]byte(sql)[:len(sql)-1]) + ") value(" //去掉点

	for _, v := range values {
		v = strings.Replace(v, "'", "\\'", -1) //转义
		sql += "'" + v + "',"
	}

	sql = string([]byte(sql)[:len(sql)-1]) + ")" //去掉点

	return sql
}

/*==================================================================================*/
/*==========================增删改查通用=========made=by=lucheng======================*/
/*==================================================================================*/

// 获得数据,根据sql语句,无分页
func GetDataBySql(data interface{}, sql string, args ...interface{}) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	// 完善可变长参数赋值问题
	var value []interface{}
	value = append(value, args[:]...)

	dba := DB.Raw(sql, value[:]...).Scan(data)
	num := dba.RowsAffected

	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = lib.CodeSuccess
		getinfo.Msg = lib.MsgSuccess
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

// 获得数据,根据name条件
func GetDataByName(data interface{}, name, value string) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	dba := DB.Where(name+" = ?", value).Find(data) //只要一行数据时使用 LIMIT 1,增加查询性能
	num := dba.RowsAffected

	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = lib.CodeSuccess
		getinfo.Msg = lib.MsgSuccess
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

// inner join
// 查询数据约定,表名_字段名(若有重复)
// 获得数据,根据id,两张表连接尝试
func GetDoubleTableDataById(model, data interface{}, id, table1, table2 string) interface{} {
	sql := fmt.Sprintf("select %s from `%s` inner join `%s` "+
		"on `%s`.%s_id=`%s`.id where `%s`.id=? limit 1", GetDoubleTableColumnsql(model, table1, table2), table1, table2, table1, table2, table2, table1)

	return GetDataBySql(data, sql, id)
}

// left join
// 查询数据约定,表名_字段名(若有重复)
// 获得数据,根据id,两张表连接
func GetLeftDoubleTableDataById(model, data interface{}, id, table1, table2 string) interface{} {

	sql := fmt.Sprintf("select %s from `%s` left join `%s` "+
		"on `%s`.%s_id=`%s`.id where `%s`.id=? limit 1", GetDoubleTableColumnsql(model, table1, table2), table1, table2, table1, table2, table2, table1)

	return GetDataBySql(data, sql, id)
}

// 获得数据,根据id
func GetDataById(data interface{}, id string) interface{} {
	var info interface{}
	var getinfo lib.GetInfoN

	dba := DB.First(data, id) //只要一行数据时使用 LIMIT 1,增加查询性能
	num := dba.RowsAffected

	//有数据是返回相应信息
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//统计页码等状态
		getinfo.Status = lib.CodeSuccess
		getinfo.Msg = lib.MsgSuccess
		getinfo.Data = data //数据
		info = getinfo
	}
	return info
}

// 获得数据,分页/查询,遵循一定查询规则,两张表,使用left join
// 如table2中查询,字段用table2_+"字段名",table1字段查询不变
func GetLeftDoubleTableDataBySearch(model, data interface{}, table1, table2 string, args map[string][]string) interface{} {
	//级联表的查询
	sqlnolimit, sql, clientPage, everyPage := GetDoubleSearchSql(model, table1, table2, args)
	sql = strings.Replace(sql, "inner join", "left join", -1)
	sqlnolimit = strings.Replace(sqlnolimit, "inner join", "left join", -1)

	return GetDataBySqlSearch(data, sql, sqlnolimit, clientPage, everyPage)
}

// 获得数据,分页/查询,遵循一定查询规则,两张表,默认inner join
// 如table2中查询,字段用table2_+"字段名",table1字段查询不变
func GetDoubleTableDataBySearch(model, data interface{}, table1, table2 string, args map[string][]string) interface{} {
	//级联表的查询以及
	sqlnolimit, sql, clientPage, everyPage := GetDoubleSearchSql(model, table1, table2, args)

	return GetDataBySqlSearch(data, sql, sqlnolimit, clientPage, everyPage)
}

// 获得数据,根据sql语句,分页
// args : sql参数'？'
// sql, sqlnolimit args 相同, 共用args
func GetDataBySqlSearch(data interface{}, sql, sqlnolimit string, clientPage, everyPage int64, args...interface{}) interface{} {
	var info interface{}
	//dest = dest.(reflect.TypeOf(dest).Elem())//type != type?
	var getinfo lib.GetInfo

	// 完善可变长参数赋值问题
	var value []interface{}
	value = append(value, args[:]...)

	dba := DB.Raw(sqlnolimit, value[:]...).Scan(&getinfo.Pager)
	num := getinfo.Pager.SumPage
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapNoResult
	} else {
		//DB.Debug().Find(&dest)
		dba = DB.Raw(sql, value[:]...).Scan(data)
		if dba.Error != nil {
			info = lib.GetSqlError(dba.Error.Error())
			return info
		}
		getinfo.Status = lib.CodeSuccess
		getinfo.Msg = lib.MsgSuccess
		getinfo.Data = data //数据

		switch {
		case strings.Contains(sql, "limit"):
			//统计页码等状态
			getinfo.Pager.ClientPage = clientPage
			getinfo.Pager.EveryPage = everyPage
			info = getinfo
		default:
			info = getinfo.GetInfoN
		}
	}
	return info
}

// 获得数据,分页/查询
func GetDataBySearch(model, data interface{}, tablename string, args map[string][]string) interface{} {

	sqlnolimit, sql, clientPage, everyPage := GetSearchSql(model, tablename, args)

	return GetDataBySqlSearch(data, sql, sqlnolimit, clientPage, everyPage)
}

// 删除通用,任意参数
func DeleteDataByName(tablename string, key, value string) interface{} {
	var info interface{}
	sql := fmt.Sprintf("delete from `%s` where %s=?", tablename, key)

	//sql = string([]byte(sql)[:len(sql)-5])
	dba := DB.Exec(sql, value)

	num := dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapDelete
	}
	return info
}

// 修改数据,通用
func UpdateDataBySql(sql string, args ...interface{}) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	var value []interface{}
	value = append(value, args[:]...)

	dba := DB.Exec(sql, value[:]...)
	num = dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapUpdate
	}
	return info
}

// 修改数据,通用
func UpdateData(tablename string, args map[string][]string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	sql, _ := GetUpdateSql(tablename, args)

	dba := DB.Exec(sql)
	num = dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapUpdate
	}
	return info
}

// 结合struct修改
func UpdateStructData(data interface{}) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	dba := DB.Save(data)
	num = dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapUpdate
	}
	return info
}

// 创建数据,通用
func CreateData(tablename string, args map[string][]string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	sql := GetInsertSql(tablename, args)
	dba := DB.Exec(sql)
	num = dba.RowsAffected

	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapError
	} else {
		info = lib.MapCreate
	}
	return info
}

// 创建数据,通用
// 返回id,事务,慎用
// 业务少可用
func CreateDataResID(tablename string, args map[string][]string) interface{} {
	var info interface{}
	//开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	sql := GetInsertSql(tablename, args)
	dba := tx.Exec(sql)
	num := dba.RowsAffected

	var value Value
	tx.Raw("select max(id) as value from " + tablename).Scan(&value)

	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapError
	} else {
		info = map[string]interface{}{"status": 201, "msg": "创建成功", "id": value.Value}
	}

	if tx.Error != nil {
		tx.Rollback()
	}

	tx.Commit()
	return info
}

// 结合struct创建
func CreateStructData(data interface{}) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	dba := DB.Create(data)
	num = dba.RowsAffected

	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapError
	} else {
		info = lib.MapCreate
	}
	return info
}

// select检查是否存在
func ValidateData(sql string) interface{} {
	var info interface{}
	var num int64 //返回影响的行数

	var ve Value
	dba := DB.Raw(sql).Scan(&ve)
	num = dba.RowsAffected

	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapError
	} else {
		info = lib.MapValidate
	}
	return info
}