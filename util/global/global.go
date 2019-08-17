package global

import der "github.com/dreamlu/go-tool"

// init param
var DBTool = &der.DBTool{}
var Config = &der.Config{}
//var Log = &der.Log{}

func init()  {
	DBTool.NewDBTool()
	DBTool.DB.LogMode(true)
	Config.NewConfig()
	//Log.DefaultFileLog()
}
