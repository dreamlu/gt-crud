// author: dreamlu
package order

import (
	"demo/models/order"
	"demo/util/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

var p order.Order

//用户信息分页
func GetOrderBySearch(u *gin.Context) {
	u.JSON(http.StatusOK, result.ResPager(p.GetMoreBySearch(result.ToCMap(u))))
}
