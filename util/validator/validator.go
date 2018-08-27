package validator

import (
	"deercoder-gin/util/lib"
	"regexp"
	"strings"
	)

/*常用验证*/
//正确返回nil,字符串验证,验证规则
func CheckRegular(value string) interface{} {
	var info interface{}

	switch {
	//手机号
	case strings.Contains(value, "phone"):
		if b, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, value); !b {
			return lib.MapPhone
		}
	//邮箱
	case strings.Contains(value, "email"):
		if b, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, value); !b {
			return lib.MapEmail
		}
	}
	return info
}

//正确返回nil,必填验证,验证规则
//value:值,name:参数名,validator:验证方式
func CheckValidator(value,name,validator string) interface{} {
	var info interface{}

	switch validator{
	//必填项
	case "required":
		if value == ""{
			info = lib.GetMapDataError(lib.CodeRequired,name+"-->必填项")
		}
	}
	return info
}

//map[string][]interface{}
func CheckValueMapArray(args map[string][]string) interface{}{
	var info interface{}
	for _,v := range args{
		info = CheckRegular(v[0])
	}
	return info
}