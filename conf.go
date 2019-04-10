// author:  dreamlu
package deercoder

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

// get path file params
func Config(key, path string) string {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Println("[CONFIG ERROR]: config file not exit")
		os.Exit(1)
	}
	return cfg.Section("").Key(key).String()
}

// default get conf/app.conf
func GetConfigValue(key string) string {

	return Config(key, "conf/app.conf")
}

// devMode file config
// devMode like dev/prod, and must has app-dev.conf/app-prod.conf
func GetDevModeConfig(key string) string {

	// dev mode
	devMode := GetConfigValue("devMode")
	if devMode == ""{
		return GetConfigValue(key)
	}

	// mode config file
	// get devMode value
	value := Config(key, "conf/app-"+devMode+".conf")
	if "" == value{
		// read the default config
		value = GetConfigValue(key)
	}
	return value
}
