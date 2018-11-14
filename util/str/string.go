package str

import "deercoder-gin/conf"

/*变量常量存储*/

var Check1 = "未认证"
var Check2 = "审核中"
var Check3 = "已认证"
var Check4 = "未通过"

/*最大上传文件大小*/
var MaxUploadMemory int64

//页码,每页数量
var ClientPageStr = conf.GetConfigValue("clientPage") //默认第1页
var EveryPageStr = conf.GetConfigValue("everyPage")   //默认10页


