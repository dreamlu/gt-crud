// @author  dreamlu
package models

import (
	"github.com/dreamlu/go-tool"
)

/*userinfo model*/
type Userinfo struct {
	ID         int64              `json:"id"`
	UserID     int64              `json:"user_id"`
	Userinfo   string             `json:"userinfo"`
	Updatetime der.JsonTime `json:"updatetime"` //maybe you like util.JsonDate
}

/*detail userinfo and user model*/
type UserinfoDe struct {
	ID         int64              `json:"id"`
	UserID     int64              `json:"user_id"`
	UserName   string             `json:"user_name"` //table `user` + `user`.name
	Userinfo   string             `json:"userinfo"`
	Updatetime der.JsonTime `json:"updatetime"`
}

// old, please see order.go
// user detail info
// include table `user` and `userinfo` data
// maybe you need to build detail info like model UserinfoBK
func GetUserInfoBySearch(args map[string][]string) interface{} {

	var userdetail []UserinfoDe
	return der.GetDoubleTableDataBySearch(UserinfoDe{}, &userdetail, "userinfo", "user", args)
}