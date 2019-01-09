### deercoder-gin
deercoder-gin is a common api development util, to help you write web api  

it can be most golang api development demo  
fragment : gin + gorm + mysql + casbin + go-ini  

##### principle：  
1.abstract func    
2.golang reflect interface{}  

##### feature：  

1.return json data  
2.one table crud,limit  
3.add two table operation(...waiting for being beeter)  
4.add basic info about website  
5.select replace * by reflect 
6.add logger for gorm 
7.add permission(user-->group(role)-->menu(permission) 
8.add validator  
9.add mysql remote connection  

##### use demo 
- create
```go
// create user
func CreateUser(args map[string][]string) interface{} {

	return db.CreateData("user", args)
}
```
- update

```go
// update user
func UpdateUser(args map[string][]string) interface{} {

	return db.UpdateData("user", args)
}
```
- delete
```go
// delete user, by id
func DeleteUserByid(id string) interface{} {

	return db.DeleteDataByName("user", "id", id)
}
```

- get information, limit and search
```go
// get user, limit and search
// clientPage(page num) 1, everyPage(every page num) 10 default
func GetUserBySearch(args map[string][]string) interface{} {
	
	var users []*User
	return db.GetDataBySearch(User{}, &users, "user", args) // anonymous User{}
}
```
- return json
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
- get information by id
```go
// get user, by id
func GetUserById(id string) interface{} {

	db.DB.AutoMigrate(&User{})
	var user = User{}
	return db.GetDataById(&user, id)
}
```
- get some data from two tables
```go
// user detail info
// include table `user` and `userinfo` data
// maybe you need to build detail info like model UserinfoBK
func GetUserinfoBySearch(args map[string][]string) interface{} { //inner join 

	var userdetail []UserinfoDe
	return db.GetDoubleTableDataBySearch(UserinfoDe{}, &userdetail, "userinfo", "user", args)
}
```