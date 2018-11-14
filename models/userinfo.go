package models

import (
	"deercoder-gin/util"
	"deercoder-gin/util/db"
)

/*userinfo model*/
type Userinfo struct {
	ID         int64         `json:"id"`
	UserID     int64         `json:"user_id"`
	Userinfo   string        `json:"userinfo"`
	Updatetime util.JsonTime `json:"updatetime"` //maybe you like util.JsonDate
}

/*detail userinfo and user model*/
type UserinfoDe struct {
	ID         int64         `json:"id"`
	UserID     int64         `json:"user_id"`
	UserName   string        `json:"user_name"` //table `user` + `user`.name
	Userinfo   string        `json:"userinfo"`
	Updatetime util.JsonTime `json:"updatetime"`
}

// user detail info
// include table `user` and `userinfo` data
// maybe you need to build detail info like model UserinfoBK
func GetUserinfoBySearch(args map[string][]string) interface{} {

	var userdetail []UserinfoDe
	return db.GetDoubleTableDataBySearch(UserinfoDe{}, &userdetail, "userinfo", "user", args)
}
