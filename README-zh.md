deercoder-gin 是一个通用的api快速开发框架示例

框架构成：gin + gorm + mysql + casbin + go-ini

##### 通用原理：

1.封装  
2.golang reflect interface{}  

##### 特点:

1.返回json数据  
2.一张表的增删改查以及分页  
3.增加多张表连接操作(...waiting for being beeter)  
4.增加网站基本信息接口  
5.select * 的优化(反射替换*为具体字段名)  
6.增加日志(定期清理待完善...)  
7.增加权限(用户-组(角色)-权限(菜单))

##### 使用示例  
- 新增
```go
// create user
func CreateUser(args map[string][]string) interface{} {

	return db.CreateData("user", args)
}
```
- 修改

```go
// update user
func UpdateUser(args map[string][]string) interface{} {

	return db.UpdateData("user", args)
}
```
- 删除
```go
// delete user, by id
func DeleteUserByid(id string) interface{} {

	return db.DeleteDataByName("user", "id", id)
}
```

- 分页,搜索二合一
```go
// get user, limit and search
// clientPage(页码) 1, everyPage(当前页) 10 default
func GetUserBySearch(args map[string][]string) interface{} {
	
	var users []*User
	return db.GetDataBySearch(User{}, &users, "user", args) //匿名User{}
}
```
- 返回json
```go
{
    "status":200,
    "msg":"请求成功",
    "data":[
        {
            "id":8,
            "name":"梦",
            "createtime":"2018-08-08 00:00:00"
        },
        {
            "id":7,
            "name":"梦",
            "createtime":"2018-08-08 00:00:00"
        }
    ],
    "pager":{
        "clientpage":1,
        "sumpage":6,
        "everypage":2
    }
}
```
- 根据id获得信息
```go
// get user, by id
func GetUserById(id string) interface{} {

	db.DB.AutoMigrate(&User{})
	var user = User{}
	return db.GetDataById(&user, id)
}
```
- 表连接,分页搜索二合一
```go
// user detail info
// include table `user` and `userinfo` data
// maybe you need to build detail info like model UserinfoBK
func GetUserinfoBySearch(args map[string][]string) interface{} {

	var userdetail []UserinfoDe
	return db.GetDoubleTableDataBySearch(UserinfoDe{}, &userdetail, "userinfo", "user", args)
}
```