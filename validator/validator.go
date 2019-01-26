// author:  dreamlu
package validator

import (
	"errors"
	"fmt"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"regexp"
	"strconv"
	"strings"
)

type Validator struct {
	data map[string][]string //要校验的数据字典
	rule map[string]*vRule //规则列表，key为字段名
}

type vRule struct {
	vr       ValidateRuler
	//required bool
}

//校验规则接口，支持自定义规则
type ValidateRuler interface {
	Check(data string) error
}

//内置规则结构，实现ValidateRuler接口
type normalRule struct {
	key    string
	trans  string //翻译后的字段名
	rule   string
	params string
}

//创建校验器对象
func NewValidator(data map[string][]string) *Validator {
	v := &Validator{data: data}
	v.rule = make(map[string]*vRule)
	return v
}

//添加内置的校验规则(同一个key只能有一条规则，重复添加会覆盖)
func (this *Validator) AddRule(key string, trans, rule string, params string) {
	nr := &normalRule{key, trans, rule, params}
	this.rule[key] = &vRule{nr}//, true} //默认required = true
}

//框架不可能包括所有的规则，为了满足不同应用的需要，除了内置规则外，需同时支持自定义规则的添加。
//go不支持重载，所以定义一个新方法来添加自定义规则（使用ValidateRuler interface参数）
func (this *Validator) AddExtRule(key string, rule ValidateRuler, required ...bool) {
	this.rule[key] = &vRule{rule} //, true}
	//if len(required) > 0 {
	//	this.rule[key].required = required[0]
	//}
}

// 执行检查后返回信息
// trans 翻译后的字段名
func (this *Validator) CheckInfo() interface{} {
	if err := this.Check(); err != nil {
		//检查不通过，处理错误
		//fmt.Println(err)
		//return err
		for k, _ := range this.rule {
			if err[k] != nil{
				return lib.GetMapData(lib.CodeValidator,err[k].Error())
			}
		}
	}
	return nil
}

//执行检查
func (this *Validator) Check() (errs map[string]error) {

	errs = make(map[string]error)
	for k, v := range this.rule {
		data, _ := this.data[k]
		if err := v.vr.Check(data[0]); err != nil { //调用ValidateRuler接口的Check方法来检查
			errs[k] = err
		}
	}
	return errs
}

// common rule
// for string, params must number-number
//
func (this *normalRule) Check(data string) (Err error) {
	//if this.params == "" {
	//	return errors.New("rule error: params wrong of rule")
	//}
	// split rule
	// 先后规则顺序
	rules := strings.Split(this.rule,",")
	for _,v := range rules{
		switch v {
		case "len":
			//字符串，根椐params判断长度的最大值和最小值
			lg := len([]rune(data))	//fix 中英文字符数量不统一
			args := strings.Split(this.params, "-")
			min, _ := strconv.Atoi(args[0])
			max, _ := strconv.Atoi(args[1])
			switch {
			case lg < min || lg > max:
				return errors.New(fmt.Sprintln(this.trans, "长度在", min, "与", max, "之间"))
			default:
				return nil
			}
		case "required":
			if data == "" {
				return errors.New(fmt.Sprintln(this.trans, "为必填项"))
			}
		case "phone":
			if b, _ := regexp.MatchString(`^1[2-9]\d{9}$`, data); !b {
				return errors.New(fmt.Sprintln("手机号格式非法"))
			}
		case "email":
			if b, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, data); !b {
				return errors.New(fmt.Sprintln("邮箱格式非法"))
			}
		case "number":
			//判断是否整数数字
			//判断最大值和最小值是否在params指定的范围
		case "list":
			//判断值是否在params指定的列表
		case "regular":
			//是否符合正则表达式
		default:
			return errors.New(fmt.Sprintf("rule error: not support of rule=%s", this.rule))
		}
	}
	return nil
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