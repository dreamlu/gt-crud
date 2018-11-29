package validator

import (
	"errors"
	"fmt"
	"github.com/Dreamlu/deercoder-gin/util/lib"
	"regexp"
	"strings"
)

type Validator struct {
	data map[string]string //要校验的数据字典
	rule map[string]*vRule //规则列表，key为字段名
}

type vRule struct {
	vr       ValidateRuler
	required bool
}

//校验规则接口，支持自定义规则
type ValidateRuler interface {
	Check(data string) error
}

//内置规则结构，实现ValidateRuler接口
type normalRule struct {
	key    string
	rule   string
	params string
}

//创建校验器对象
func NewValidator(data map[string]string) *Validator {
	v := &Validator{data: data}
	v.rule = make(map[string]*vRule)
	return v
}

//添加内置的校验规则(同一个key只能有一条规则，重复添加会覆盖)
func (this *Validator) AddRule(key string, rule string, params string, required ...bool) {
	nr := &normalRule{key, rule, params}
	this.rule[key] = &vRule{nr, true} //默认required = true
	if len(required) > 0 {
		this.rule[key].required = required[0]
	}
}

//框架不可能包括所有的规则，为了满足不同应用的需要，除了内置规则外，需同时支持自定义规则的添加。
//go不支持重载，所以定义一个新方法来添加自定义规则（使用ValidateRuler interface参数）
func (this *Validator) AddExtRule(key string, rule ValidateRuler, required ...bool) {
	this.rule[key] = &vRule{rule, true}
	if len(required) > 0 {
		this.rule[key].required = required[0]
	}
}

//执行检查
func (this *Validator) Check() (errs map[string]error) {
	errs = make(map[string]error)
	for k, v := range this.rule {
		data, exists := this.data[k]
		if !exists { //无值
			if v.required { //如果必填，报错
				errs[k] = errors.New("data error: required field miss")
			}
		} else { //有值判断规则
			if err := v.vr.Check(data); err != nil { //调用ValidateRuler接口的Check方法来检查
				errs[k] = err
			}
		}
	}
	return errs
}

func (this *normalRule) Check(data string) (Err error) {
	if this.params == "" {
		Err = errors.New("rule error: params wrong of rule")
		return
	}
	switch this.rule {
	case "string":
		Err = errors.New("范围判断")
		//字符串，根椐params判断长度的最大值和最小值
	case "number":
		//判断是否整数数字
		//判断最大值和最小值是否在params指定的范围
	case "list":
		//判断值是否在params指定的列表
	case "regular":
		//是否符合正则表达式
	default:
		Err = errors.New(fmt.Sprintf("rule error: not support of rule=%s", this.rule))
	}
	return
}

type myRuler struct {
}

//添加Check方法，实现ValidateRuler 接口
func (this *myRuler) Check(data string) (Err error) {
	//判断data是否符合规则
	return
}

//添加规则
//validator.AddExtRule("name", &myRuler{})

/*常用验证*/
//正确返回nil,字符串验证,验证规则
func CheckRegular(value string) interface{} {
	//反射得到变量名
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
func CheckValidator(value, name, validator string) interface{} {
	var info interface{}

	switch validator {
	//必填项
	case "required":
		if value == "" {
			info = lib.GetMapData(lib.CodeRequired, name+"-->必填项")
		}
	}
	return info
}

//map[string][]interface{}
func CheckValueMapArray(args map[string][]string) interface{} {
	var info interface{}
	for _, v := range args {
		info = CheckRegular(v[0])
	}
	return info
}
