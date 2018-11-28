package deercoder

import (
	"fmt"
	"testing"
)

var (
	Host = "127.0.0.1:6379"
	Password = ""
	Database = 0
	MaxOpenConns = 0 // max number of connections
	MaxIdleConns = 0 // 最大的空闲连接数
)

var R *ConnPool

func TestRedis(t *testing.T) {
	R.SetString("test","testvalue")
	value,_ := R.GetString("test")
	fmt.Println(value)
}

// in test, init no use
func init()  {
	R = InitRedisPool(Host, Password, Database, MaxOpenConns, MaxIdleConns)
}
