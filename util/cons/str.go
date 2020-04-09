package cons

import "github.com/dreamlu/gt"

var (
	DevMode = gt.Configger().GetString("app.devMode")
	Version = gt.Configger().GetString("app.version")
)
