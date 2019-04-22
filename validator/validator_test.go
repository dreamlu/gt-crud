// @author  dreamlu
package validator

import (
	"fmt"
	"testing"
)

// test validator
func TestValidator(t *testing.T) {
	var maps = make(map[string][]string)
	maps["name"] = append(maps["name"], "梦")
	val := NewValidator(maps) //将要检查的数据字典传入，生成Validator对象
	val.AddRule("name", "用户名","required,len","2-5") //对字段name添加规则： 2-5个字符长度，必填
	//val.AddRule("sport","list","football,swim",false) //对字段sport添加规则： 值需在列表中（football，swim),非必填

	info := val.CheckInfo()
	switch info {
	case nil:
		fmt.Println("pass")
	default:
		fmt.Println(info)
	}

}

