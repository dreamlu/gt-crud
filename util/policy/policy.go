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
// 1: 项目部领导: 和0一样,但是只能看不能操作
// 2: 工程领导: 只能查看大屏和修改自己的状态
// 3: 项目部人员: 任务交接,文件签认,文件审批,影像资料上传,pdf签字
// 4: (工程下)施工部人员: 资料自动填写
// 5: 影像资料管理员: 影像资料管理(文件夹创建,文件删除等...)
// 6: 天窗负责人
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

		// 2: 工程领导
		role = 2
		if containPaths(path, "/engine/lead/update", "/engine/lead/get") {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
		}

		// 3: 项目部人员
		role = 3
		if containPaths(path, "/project/task",
			"/project/signature",
			"/project/approve",
			"/project/resource/search",
			"/project/resource/create",
			"/project/work/order",
			"/office/pdf",
		) {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
		}
		// 4: 施工部人员
		role = 4
		if containPaths(path, "/office",
			"/engine/issue",
		) && !containPaths(path, "/office/pdf") {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
		}
		// 5: 影像资料管理员
		role = 5
		if containPaths(path, "/project/resource") {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
		}
		// 6: 天窗负责人
		role = 6
		if containPaths(path, "/construct/dormer") {
			rolePs = append(rolePs, fmt.Sprintf(format, role, path, "*"))
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
