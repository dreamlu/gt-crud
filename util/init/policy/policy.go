package policy

import (
	authz2 "demo/routers/authz"
	"demo/routers/routelist"
	"demo/util/cons"
	"fmt"
	"net/http"
	"strings"
)

// InitPolicy 权限策略初始化
// 定义角色策略
// role:
// -1: admin直接返回不验证,所有权限
// 0: 项目部总调度: 所有权限除了admin相关
func InitPolicy() {

	var (
		rolePs []string
		format = "%d,%s,%s"
	)

	// 权限角色整理
	for route, methods := range routelist.RouteList.List {
		//log.Info(route)
		var (
			path = strings.TrimPrefix(route, cons.Prefix)
			role = 0
		)
		// DELETE方法
		path = strings.TrimSuffix(path, "/:id")
		if !containPaths(path, "/admin") {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
		}

		role = 1
		if !containPaths(path, "/admin") && methods[0] == http.MethodGet {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, http.MethodGet))
		}
	}

	// 添加策略
	for _, rolePolicy := range rolePs {
		rs := strings.Split(rolePolicy, ",")
		authz2.AddPolicy(rs...)
	}
	authz2.SavePolicy()
}

func containPaths(path string, conditions ...string) bool {
	for _, v := range conditions {
		if strings.Contains(path, v) {
			return true
		}
	}
	return false
}
