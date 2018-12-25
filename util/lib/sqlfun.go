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
func GetKeySql(sql, sqlnolimit string, key string, model interface{}, alias string) (sqlkey, sqlnolimitkey string) {

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

// 两张表， 表1包含表2 id
// key search sql
func GetDoubleKeySql(sql, sqlnolimit string, key string, model interface{}, tablel, table2 string) (sqlkey, sqlnolimitkey string) {

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
			case !strings.HasSuffix(tag, "id") && !strings.HasSuffix(tag, "date") && !strings.HasSuffix(tag, "time"):
				// 表2
				if strings.Contains(tag, table2+"_") && !strings.Contains(tag, table2+"_id") {
					sql += "`"+ table2 + "`.`" + string([]byte(tag)[len(table2)+1:]) + "` like binary '%" + key + "%' or "
					sqlnolimit += "`"+ table2 + "`.`" + string([]byte(tag)[len(table2)+1:]) + "` like binary '%" + key + "%' or "
					continue
				}

				// 表1
				sql += "`" + tablel + "`.`" + tag + "` like binary '%" + key + "%' or "
				sqlnolimit += "`" + tablel + "`.`" + tag + "` like binary '%" + key + "%' or "
			}

		}
		sql = string([]byte(sql)[:len(sql)-4]) + ") and "
		sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4]) + ") and "
	}
	//sql = string([]byte(sql)[:len(sql)-4])
	//sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4])
	return sql, sqlnolimit
}
