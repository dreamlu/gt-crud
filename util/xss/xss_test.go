package xss

import (
	"fmt"
	"testing"
)

// test xss
func TestXss(t *testing.T)  {
	var maps = make(map[string][]string)
	maps["name"] = append(maps["name"], "梦 '< and 1=1 \" --")
	fmt.Println(XssMap(maps))
}
