package client

import (
	token2 "demo/util/models/token"
	"demo/util/result"
	"demo/util/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// token
func Token(u *gin.Context) {
	client_id := u.Query("id")
	if client_id == "" {
		u.JSON(http.StatusOK, result.GetError("id不能为空"))
		return
	}
	id, err := strconv.ParseUint(client_id, 10, 64)
	if err != nil {
		u.JSON(http.StatusOK, result.Res(err))
		return
	}
	tk := token.GetToken(token2.TokenModel{
		ID:   id,
		Typ:  0,
		Role: nil,
		G:    "",
	})

	u.JSON(http.StatusOK, result.MapSuccess.Add("token", tk))
}
