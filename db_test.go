package deercoder

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Createtime JsonDate `json:"createtime"`
}

func TestDB(t *testing.T) {

	var user = User{
		Name:"测试",
		Createtime:JsonDate(time.Now()),
	}

	ss := CreateStructData(&user)
	fmt.Println(ss)

	user.ID = 8
	ss = UpdateStructData(&user)
	fmt.Println(ss)
}
