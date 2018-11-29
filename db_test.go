package deercoder

import (
	"fmt"
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