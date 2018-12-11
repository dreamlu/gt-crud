package deercoder

import (
	"fmt"
	"github.com/Dreamlu/deercoder-gin"
	"log"
	"strconv"
	"strings"
	"testing"
)

type User struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name" validate:"len=20"`
	Createtime JsonDate `json:"createtime"`
}

func TestDB(t *testing.T) {

	var user = User{
		Name: "测试",
		//Createtime:JsonDate(time.Now()),
	}

	ss := CreateStructData(&user)
	fmt.Println(ss)

	user.ID = 8
	ss = UpdateStructData(&user)
	fmt.Println(ss)
}

// 通用分页测试
// 如：
func TestSqlSearch(t *testing.T) {
	sql := fmt.Sprintln(`select %s from userinfo a inner join xx on a.xx_id=b.id
		where 1=1 and `,GetSqlColumnsSql(User{}))
	sqlnolimit := `select
		count(distinct a.id) as sum_page	
		from venuepricets a inner join venue b on a.venue_id=b.id inner join venuer c on b.venuer_id=c.id
		where 1=1 and `
	var ui []Userinfo

	//页码,每页数量
	clientPageStr := GetConfigValue("clientPage") //默认第1页
	everyPageStr := GetConfigValue("everyPage")   //默认10页

	//可定制
	//args map[string][]string
	//look deercoder-gin/demo
	for k, v := range args {
		if k == "clientPage" {
			clientPageStr = v[0]
			continue
		}
		if k == "everyPage" {
			everyPageStr = v[0]
			continue
		}
		if v[0] == "" { //条件值为空,舍弃
			continue
		}
		v[0] = strings.Replace(v[0], "'", "\\'", -1) //转义
		sql += "a." + k + " = '" + v[0] + "' and "
		sqlnolimit += "a." + k + " = '" + v[0] + "' and "
	}

	clientPage, _ := strconv.ParseInt(clientPageStr, 10, 64)
	everyPage, _ := strconv.ParseInt(everyPageStr, 10, 64)

	sql = string([]byte(sql)[:len(sql)-4]) //去and
	sqlnolimit = string([]byte(sqlnolimit)[:len(sqlnolimit)-4])
	sql += "order by id desc limit " + strconv.FormatInt((clientPage-1)*everyPage, 10) + "," + everyPageStr
	log.Println(GetDataBySqlSearch(&ui, sql, sqlnolimit, clientPage, everyPage))
}