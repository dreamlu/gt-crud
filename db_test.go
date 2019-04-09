// author:  dreamlu
package deercoder

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type User struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Createtime JsonDate `json:"createtime"`
}

type UserInfo struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`   //用户id
	UserName string `json:"user_name"` //用户名
	Userinfo string `json:"userinfo"`
}

// order
type Order struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`     // user id
	ServiceID  int64 `json:"service_id"`  // service table id
	CreateTime int64 `json:"create_time"` // createtime
}

// order detail
type OrderD struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`      // user id
	UserName    string `json:"user_name"`    // user table column name
	ServiceID   int64  `json:"service_id"`   // service table id
	ServiceName string `json:"service_name"` // service table column `name`
	Createtime  JsonTime  `json:"createtime"`   // createtime
}

func TestDB(t *testing.T) {

	var user = User{
		Name: "测试",
		//Createtime:JsonDate(time.Now()),
	}

	ss := CreateStructData(&user)
	fmt.Println(ss)

	user.ID = 8 //0
	ss = UpdateStructData(&user)
	fmt.Println(ss)
}

// 通用分页测试
// 如：
func TestSqlSearch(t *testing.T) {
	sql := fmt.Sprintf(`select a.id,a.user_id,a.userinfo,b.name as user_name from userinfo a inner join user b on a.user_id=b.id where 1=1 and `)
	sqlnolimit := `select
		count(distinct a.id) as sum_page
		from userinfo a inner join user b on a.user_id=b.id
		where 1=1 and `
	var ui []UserInfo

	//页码,每页数量
	clientPageStr := GetDevModeConfig("clientPage") //默认第1页
	everyPageStr := GetDevModeConfig("everyPage")   //默认10页

	//可定制
	//args map[string][]string
	//look deercoder-gin/demo
	//args is url.values
	//for k, v := range args {
	//	if k == "clientPage" {
	//		clientPageStr = v[0]
	//		continue
	//	}
	//	if k == "everyPage" {
	//		everyPageStr = v[0]
	//		continue
	//	}
	//	if v[0] == "" { //条件值为空,舍弃
	//		continue
	//	}
	//	v[0] = strings.Replace(v[0], "'", "\\'", -1) //转义
	//	sql += "a." + k + " = '" + v[0] + "' and "
	//	sqlnolimit += "a." + k + " = '" + v[0] + "' and "
	//}

	clientPage, _ := strconv.ParseInt(clientPageStr, 10, 64)
	everyPage, _ := strconv.ParseInt(everyPageStr, 10, 64)

	sql = string([]byte(sql)[:len(sql)-4]) //去and
	sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4])
	sql += "order by a.id desc limit " + strconv.FormatInt((clientPage-1)*everyPage, 10) + "," + everyPageStr
	log.Println(GetDataBySQLSearch(&ui, sql, sqlnolimit, clientPage, everyPage))
	log.Println(ui)
}

// 常用分页测试(两张表)
// 如:
func TestSqlSearchV2(t *testing.T) {
	//var ui []UserInfo
	//
	////args map[string][]string
	////look deercoder-gin/demo
	////args is url.values
	//log.Println(GetDoubleTableDataBySearch(UserInfo{},&ui, "userinfo", "user", args))
	//log.Println(ui)
}

// select 数据存在验证
func TestValidateData(t *testing.T) {
	sql := "select *from `user` where id=2"
	ss := ValidateSQL(sql)
	log.Println(ss)
}

// 分页搜索中key测试
func TestGetSearchSql(t *testing.T) {

	var args = make(map[string][]string)
	args["key"] = append(args["key"], "梦 嘿,伙计")
	sqlnolimit, sql, _, _ , _:= GetSearchSQL(User{}, "user", args)
	log.Println("SQLNOLIMIT:", sqlnolimit, "\nSQL:", sql)

	// 两张表，待重新测试
	sqlnolimit, sql, _, _ = GetDoubleSearchSql(UserInfo{}, "userinfo", "user", args)
	log.Println("SQLNOLIMIT==>2:", sqlnolimit, "\nSQL==>2:", sql)

}

// 通用sql以及参数
func TestGetDataBySql(t *testing.T) {
	var sql = "select id,name,createtime from `user` where id = ? and name = ?"

	var user User
	GetDataBySQL(&user, sql, "1", "梦")
	fmt.Println(user)

	//DB.Raw(sql, []interface{}{1, "梦"}[:]...).Scan(&user)
	//fmt.Println(user)
}

// 通用增删该查测试
func TestCrud(t *testing.T) {
	var args = make(map[string][]string)
	args["name"] = append(args["name"], "梦")

	// var crud DbCrud
	// must use AutoMigrate
	// get by id
	DB.AutoMigrate(&User{})
	var user User
	var db = DbCrud{
		Table:     "user",
		ModelData: &user,
	}
	info := db.GetByID("1")
	log.Println(info, "\n[User Info]:", user)

	// get by search
	var users []*User
	db = DbCrud{
		Table:     "user",
		Model:     User{},
		ModelData: &users,
	}
	args["name"][0] = "梦4"
	db.GetBySearch(args)
	log.Println("\n[User Info]:", users)

	// delete
	info = db.Delete("12")
	log.Println(info)

	// update
	args["id"] = append(args["id"], "4")
	args["name"][0] = "梦4"
	info = db.Update(args)
	log.Println(info)

	// create
	//var args2 = make(map[string][]string)
	//args2["name"] = append(args2["name"],"梦c")
	////db  = DbCrud{"user", nil,&user}
	//info = db.Create(args2)
	//log.Println(info)

	// get more search
	var params = make(map[string][]string)
	var or []*OrderD
	db = DbCrud{
		InnerTables: []string{"order", "user"},
		LeftTables:  []string{"service"},
		Model:       OrderD{},
		ModelData:   &or,
	}
	db.GetMoreBySearch(params)
	log.Println("\n[User Info]:", or[0])
}

// 通用增删改查sql测试
func TestCrudSQL(t *testing.T) {
	var db = DbCrud{}
	sql := "update `user` set name=? where id=?"
	log.Println("[Info]:", db.UpdateBySQL(sql,"梦sql", 1))
}