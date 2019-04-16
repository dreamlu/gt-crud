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
| select `*`优化<br>反射替换*为具体字段名 |
| 优化自定义gorm日志<br>存储错误sql以及相关error |  
| 增加权限<br>用户-组(角色)-权限(菜单)(待优化) |
| 增加参数验证 |
| 增加mysql远程连接 |
| 增加多表key模糊搜索 |
| session(cookie/redis) |
| 更多数据库支持(待完善) |
| conf/app.conf 多开发模式支持 |
| 请求方式json/form data(待测试) |

##### 使用  
- [安装使用](#安装使用)
- [API 使用](#api-examples)
    - [Create](#create)
    - [Update](#update)
    - [Delete](#delete)
    - [GetBySearch](#getbysearch)
    - [GetByID](#getbyid) 
    - [GetMoreBySearch](#getmorebysearch)
    - [GetDataBySQL](#getdatabysql)
    - [GetDataBySearchSQL](#getdatabysearchsql)
    - [DeleteBySQL](#deletebysql)
    - [UpdateBySQL](#updatebysql)
    - [CreateBySQL](#createbysql)
    - [GetDevModeConfig](#getdevmodeconfig)
    

#### 安装使用  
1.下载demo  
2.全局替换demo为你的项目名  
3.go build  

### API Examples  
#### Create
```go
// dbcrud
var db = deercoder.DbCrud{
	Model: User{},		// model
	Table:"user",		// table name
}

// create user
func (c *User)Create(params map[string][]string) interface{} {

	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	return db.Create(params)
}
```

#### Update
```go
// update user
func (c *User)Update(params map[string][]string) interface{} {

	return db.Update(params)
}
```

#### Delete
```go
// delete user, by id
func (c *User)Delete(id string) interface{} {

	return db.Delete(id)
}
```

#### GetBySearch
```go
// get user, limit and search
// clientPage 1, everyPage 10 default
func (c *User)GetBySearch(params map[string][]string) interface{} {
	var users []*User
	db.ModelData = &users
	return db.GetBySearch(params)
}
```

#### GetByID
```go
// get user, by id
func (c *User)GetByID(id string) interface{} {

	var user User	// not use *User
	db.ModelData = &user
	return db.GetByID(id)
}
```

#### GetMoreBySearch
```go
// get order, limit and search
// clientPage 1, everyPage 10 default
func (c *Order) GetMoreBySearch(params map[string][]string) interface{} {
	var or []*OrderD
	db = deercoder.DbCrud{
		InnerTables: []string{"order", "user"}, // inner join tables, 'order' must the first table
		LeftTables:  []string{"service"},       // left join tables
		Model:       OrderD{},                  // order model
		ModelData:   &or,                       // model value
	}
	return db.GetMoreBySearch(params)
}

```

#### GetDataBySQL
```go
// like UpdateBySQL
```

#### GetDataBySearchSQL
```go
// like UpdateBySQL
```

#### DeleteBySQL
```go
// like UpdateBySQL
```

#### UpdateBySQL
```go
var db = DbCrud{}
sql := "update `user` set name=? where id=?"
log.Println("[Info]:", db.UpdateBySQL(sql,"梦sql", 1))
```

#### CreateBySQL
```go
// like UpdateBySQL
```

- 多模式配置文件  
> 配置方式: conf/app.conf 中 `devMode = dev` 对应conf/app-`dev`.conf  

#### GetDevModeConfig
```go
// devMode test
// app.conf devMode = dev
// test the app-dev.conf value
func TestDevMode(t *testing.T)  {
	log.Println("config read test: ", GetDevModeConfig("db.host"))
}
```