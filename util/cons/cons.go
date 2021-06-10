package cons

import (
	"github.com/dreamlu/gt/tool/conf"
)

var (
	DevMode = conf.GetString("app.devMode")
	Version = conf.GetString("app.version")
	Prefix  = conf.GetString("app.prefix")
)
