#### 后端开发规范设计  
deercoder-gin 是一个通用的api快速开发工具  

##### 所需基础知识:  
| 知识   | 链接 | 
| ---   | ---- |
| go基础 |  [go链接](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md) |
| gin   | [gin](https://github.com/gin-gonic/gin) |
| mysql | mysql语法 |

##### 工具构成:  
| 路由 | orm  | 数据库 | 权限   | 配置    |
| --- | ---- | ----  | ------ | ------ |
| gin | gorm | mysql | casbin | go-ini | 

或许你需要下载结合[api.html](./demo/api.html)接口文档查看测试以及查看*_test.go文件
##### 通用原理：

1.封装  
2.golang reflect interface{}  

##### 特点:
| 特点 | 
| ------ |
| restful api |  
| 一张表的增删改查以及分页 |   
| 多张表连接操作 |  
| 增加网站基本信息接口 |  
| select * 的优化(反射替换*为具体字段名) |
| 优化自定义gorm日志(存储错误sql以及相关error) |  
| 增加权限(用户-组(角色)-权限(菜单))(待优化) |
| 增加参数验证 |
| 增加mysql远程连接 |
| 增加多表key模糊搜索 |
| session(cookie/redis) |
| 更多数据库支持(待完善) |

##### 使用示例  
- 新增
```go
// create user
func (c *User)Create(args map[string][]string) interface{} {

	args["createtime"] = append(args["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	return deercoder.CreateData("user", args)
}
```
- 修改

```go
// update user
func (c *User)Update(args map[string][]string) interface{} {

	return deercoder.UpdateData("user", args)
}
```
- 删除
```go
// delete user, by id
func (c *User)DeleteByid(id string) interface{} {

	return deercoder.DeleteDataByName("user", "id", id)
}
```

- 分页,搜索二合一
```go
// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User)GetBySearch(args map[string][]string) interface{} {
	//相当于注册类型,https://github.com/jinzhu/gorm/issues/857
	//db.DB.AutoMigrate(&User{})
	//var users = []*User{}
	var users []*User
	return deercoder.GetDataBySearch(User{}, &users, "user", args) //匿名User{}
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
func (c *User)GetById(id string) interface{} {

	deercoder.DB.AutoMigrate(&User{})
	var user = User{}
	return deercoder.GetDataById(&user, id)
}
```
- 表连接,分页搜索二合一
```go
// user detail info
// include table `user` and `userinfo` data
// maybe you need to build detail info like model UserinfoBK
func GetUserinfoBySearch(args map[string][]string) interface{} {//inner join 

	var userdetail []UserinfoDe
	return db.GetDoubleTableDataBySearch(UserinfoDe{}, &userdetail, "userinfo", "user", args)
}
```