// @author  dreamlu
package lib

import (
	"strings"
)

// 数据库错误过滤、转换(友好提示)
func GetSqlError(error string) (info MapData) {
	switch {
	case error == "record not found":
		info = MapNoResult
	case strings.Contains(error, "Duplicate entry"):
		//error = strings.Replace(error, "Error 1062: Duplicate entry", "", -1)
		errors := strings.Split(error, "for key ")
		//error = "已存在相同数据:" + errors[0]
		error = strings.Trim(errors[1], "'") //自定义数据库唯一约束名
		info = GetMapData(CodeText, error)
	case strings.Contains(error, "Data too long"):
		error = "存在字段范围过长"
		info = GetMapData(CodeText, error)
	default:
		info = GetMapData(CodeText, error)
	}

	return info
}
