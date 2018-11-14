package models

import (
	"deercoder-gin/util"
	"deercoder-gin/util/db"
)

/*user model*/
type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Createtime util.JsonTime `json:"createtime"`
}

//获得用户,根据id
func GetUserById(id string) interface{} {

	db.DB.AutoMigrate(&User{})
	var user = User{}
	return db.GetDataById(&user, id)
}

//获得用户,分页/查询
func GetUserBySearch(args map[string][]string) interface{} {
	//相当于注册类型,https://github.com/jinzhu/gorm/issues/857
	db.DB.AutoMigrate(&User{})
	var users = []*User{}
	return db.GetDataBySearch(User{}, &users, "user", args) //匿名User{}
}

//删除用户,根据id
func DeleteUserByid(id string) interface{} {

	return db.DeleteDataByName("user", "id", id)
}

//修改用户
func UpdateUser(args map[string][]string) interface{} {

	return db.UpdateData("user", args)
}

//创建用户
func CreateUser(args map[string][]string) interface{} {

	return db.CreateData("user", args)
}
