package deercoder

import (
	"fmt"
	"testing"
)

const ConfigPath = "conf/app.conf"

func TestConfig(t *testing.T){
	fmt.Println("config read test: ",Config("http_port",ConfigPath))
}
