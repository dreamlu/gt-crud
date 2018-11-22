package deercoder

import (
	"github.com/gin-gonic/gin"
)

// return values
// type map[string][]string
func GetUriValues(u *gin.Context) map[string][]string {
	u.Request.ParseForm()
	values := u.Request.Form //need to ParseForm
	return values
}
