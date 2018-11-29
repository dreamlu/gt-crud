package validator

import (
	"fmt"
	"testing"
)

// test size
func TestValidator(t *testing.T) {
	var maps = map[string]string{"name":"qwertyuiopasdfghjklzx"}
	val := NewValidator(maps) //将要检查的数据字典传入，生成Validator对象
	val.AddRule("name","string","2-5",true) //对字段name添加规则： 2-5个字符长度，必填
	val.AddRule("sport","list","football,swim",false) //对字段sport添加规则： 值需在列表中（football，swim),非必填

	if err:= val.Check(); err !=nil{
		//检查不通过，处理错误
		fmt.Println(err)
		return
	}
}

