package str

import "github.com/dreamlu/gt"

var (
	DevMode = gt.Configger().GetString("app.devMode")
)
