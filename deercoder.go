package deercoder

import (
	"github.com/gin-gonic/gin"
)

// main.go
//import (
//	"github.com/Dreamlu/deercoder-gin/routers"
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	gin.SetMode(gin.DebugMode)
//	r := routers.SetRouter()
//	// Listen and Server in 0.0.0.0:8080
//	r.Run(":" + GetConfigValue("http_port"))
//}


// return values
// type map[string][]string
func GetUriValues(u *gin.Context) map[string][]string {
	u.Request.ParseForm()
	values := u.Request.Form //need to ParseForm
	return values
}