package reflect

import (
	"testing"
)

type Client struct {
	Name string `gorm:"type:varchar(30);" json:"name" valid:"required,len=2-20"` // 昵称
}

func TestNew(t *testing.T) {
	//var ct Client
	t.Log(Client{})
	t.Log(New(Client{Name: "test"}))
}
