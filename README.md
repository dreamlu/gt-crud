deercoder-gin 是一个通用的api快速开发框架

包含了快速开发示例

框架构成：gin + gorm + mysql + casbin + go-ini

通用原理：

1.封装

2.golang reflect interface{}

特点:

1.返回json数据

2.一张表的增删改查以及分页

3.增加多张表连接操作(...waiting for being beeter)

4.增加网站基本信息接口

5.select * 的优化(反射替换*为具体字段名)

6.增加日志(定期清理待完善...)

7.增加权限(用户-组(角色)-权限(菜单))