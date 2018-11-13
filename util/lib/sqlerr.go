package lib

import "kpx/lib"

//数据库错误过滤、转换(友好提示)
func GetSqlError(error string) (info interface{}){
	switch error {
	case "record not found":
		info = MapNoResult
	default:
		info = lib.GetMapDataError(CodeSql,error)
	}
	return info
}
