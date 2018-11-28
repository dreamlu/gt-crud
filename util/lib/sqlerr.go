package lib

import (
	"strings"
)

//数据库错误过滤、转换(友好提示)
func GetSqlError(error string) (info interface{}) {
	switch  {
	case error == "record not found":
		info = MapNoResult
	case strings.Contains(error, "Duplicate entry"):
		error = strings.Replace(error, "Error 1062: Duplicate entry", "", -1)
		error = strings.Replace(error, "for key 'name'", "", -1)
		error = "已存在相同数据:" + error
		info = map[string]interface{}{"status":CodeText, "msg":error}
	case strings.Contains(error, "Data too long"):
		error = "存在字段范围过长"
		info = map[string]interface{}{"status":CodeText, "msg":error}
	default:
		info = GetMapData(CodeText, error)
	}

	return info
}
