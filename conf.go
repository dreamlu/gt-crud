package deercoder

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

func Config(key, path string) string{
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("config file not exit")
		os.Exit(1)
	}
	return cfg.Section("").Key(key).String()
}


