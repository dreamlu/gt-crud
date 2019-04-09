// author:  dreamlu
package deercoder

import (
	"log"
	"testing"
)

const ConfigPath = "conf/app.conf"

func TestConfig(t *testing.T){
	log.Println("config read test: ",Config("http_port",ConfigPath))
}

// devMode test
// app.conf devMode = dev
// test the app-dev.conf value
func TestDevMode(t *testing.T)  {
	log.Println("config read test: ", GetDevModeConfig("db.host"))
}
