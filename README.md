deercoder-gin 

deercoder-gin 是一个gin + gorm + gt 的使用案例  
- 特点  
1.[gin](https://github.com/gin-gonic/gin) 提供核心的请求处理(了解)  
2.[gorm](https://github.com/jinzhu/gorm) 数据库处理(了解)  
3.[gt](https://github.com/dreamlu/gt) 提供常用的操作工具类以及一些开发约定(封装了gin+gorm形成常用业务链,须知)  

- api文档参考：  
1.api.html(或者在线[api.html](https://www.eolinker.com/#/share/project/api/?groupID=-1&shareCode=pgnwpF&shareToken=$2y$10$QMWRQU4fEfGOLkZgLwGFX.UHcWaaR1Eutrh6DCG8u0XKDRwwcUv76&shareID=120217))  
2.此处单机部署开发,单机docker化参考docker目录,微服务go参考[micro-go](https://github.com/dreamlu/micro-go)  

- 数据库模型生成,[代码](./util/db/db.go)  
> 模型定义需遵循:[模型定义](https://gorm.io/zh_CN/docs/models.html)  
- 插件[代码](./util/plugin/README.md)    
> 插件提供了一些其他常见功能  

- 开箱即用  
1.针对小程序  
2.如用到支付回调和退款回调,请修改conf/中notifyUrl为你自己本地测试域名和上线域名,并根据注释掉的代码书写自己的逻辑  
3.默认多账号小程序;单账号:注释掉models/global.go中AdminCom的adminID字段,(ps:为了方便部署, 全局搜索`initApplet()`,打开注释,填入单账号的appid等参数)`  
n.更多用法参考[gt](https://github.com/dreamlu/gt)  

- 关于crud  
```go
var crud = gt.NewCrud(
	gt.Model(Client{}),
)
``` 
案例中使用是当前页面全局性的变量, 如果新增了方法, 修改了Model(),如:
`crud.Params(gt.Model(Client2{}))`
这种情况下,其他使用相同变量的crud的Model都会收到影响,解决方法如下:  
1.新增的方法中使用`gt.NewCrud(gt.Model(Client2{}))`解决  
2.可将`crud.Params(gt.Model(Client{}))`添加至每个使用的crud变量中  