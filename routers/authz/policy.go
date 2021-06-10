package authz

import "github.com/dreamlu/gt/tool/log"

// AddPolicy 添加角色策略
// role, path, method...
func AddPolicy(rpm ...string) {
	var params []interface{}
	for _, v := range rpm {
		params = append(params, v)
	}

	_, err := Enforcer.AddPolicy(params...)
	if err != nil {
		log.Error(err)
		return
	}
}

// SavePolicy 所有策略添加完再存储
func SavePolicy() {
	err := Enforcer.SavePolicy()
	if err != nil {
		log.Error(err)
		return
	}
}

// AddGroupPolicy 添加组策略
func AddGroupPolicy(group, role string) {
	_, err := Enforcer.AddGroupingPolicy(group, role)
	if err != nil {
		log.Error(err.Error())
	}
}
