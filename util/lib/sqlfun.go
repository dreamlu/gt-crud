// @author  dreamlu
package lib

import (
	"reflect"
	"strings"
)

// 根据model中表模型的json标签获取表字段
// 将select* 变为对应的字段名
func GetTags(model interface{}) (tags []string) {
	typ := reflect.TypeOf(model)
	//var buffer bytes.Buffer
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("json")
		tags = append(tags, tag)
	}
	return tags
}

// key search sql
func GetKeySQL(sql, sqlnolimit string, key string, model interface{}, alias string) (sqlkey, sqlnolimitkey string) {

	tags := GetTags(model)
	keys := strings.Split(key, " ") //空格隔开
	for _, key := range keys {
		if key == "" {
			continue
		}
		sql += "("
		sqlnolimit += "("
		for _, tag := range tags {
			switch {
			// 排除id结尾字段
			// 排除date,time结尾字段
			case !strings.HasSuffix(tag, "id") ://&& !strings.HasSuffix(tag, "date") && !strings.HasSuffix(tag, "time"):
				sql += "`" + alias + "`.`" + tag + "` like binary '%" + key + "%' or "
				sqlnolimit += "`" + alias + "`.`" + tag + "` like binary '%" + key + "%' or "
			}

		}
		sql = string([]byte(sql)[:len(sql)-4]) + ") and "
		sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4]) + ") and "
	}
	//sql = string([]byte(sql)[:len(sql)-4])
	//sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4])
	return sql, sqlnolimit
}

// 多张表, 第一个表为主表
// key search sql
// tables [table1:table1_alias]
// searModel : 搜索字段模型
func GetMoreKeySQL(sql, sqlnolimit string, key string, searModel interface{}, tables ...string) (sqlkey, sqlnolimitkey string) {

	// 搜索字段
	tags := GetTags(searModel)
	// 多表
	keys := strings.Split(key, " ") //空格隔开
	for _, key := range keys {
		if key == "" {
			continue
		}
		sql += "("
		sqlnolimit += "("
		for _, tag := range tags {
			switch {
			// 排除id结尾字段
			// 排除date,time结尾字段
			case !strings.HasSuffix(tag, "id") && !strings.HasSuffix(tag, "date") && !strings.HasSuffix(tag, "time"):

				// 多表判断
				for _, v := range tables {
					ts := strings.Split(v, ":")
					table := ts[0]
					alias := ts[1]
					if strings.Contains(tag, table+"_") && !strings.Contains(tag, table+"_id") {
						sql += "`" + alias + "`.`" + string([]byte(tag)[len(table)+1:]) + "` like binary '%" + key + "%' or "
						sqlnolimit += "`" + alias + "`.`" + string([]byte(tag)[len(table)+1:]) + "` like binary '%" + key + "%' or "
						goto into
					}
				}

				// 主表
				ts := strings.Split(tables[0], ":")
				alias := ts[1]
				sql += "`" + alias + "`.`" + tag + "` like binary '%" + key + "%' or "
				sqlnolimit += "`" + alias + "`.`" + tag + "` like binary '%" + key + "%' or "
			}
		into:
		}
		sql = string([]byte(sql)[:len(sql)-4]) + ") and "
		sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4]) + ") and "
	}
	//sql = string([]byte(sql)[:len(sql)-4])
	//sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4])
	return sql, sqlnolimit
}