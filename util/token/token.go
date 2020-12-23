package token

import (
	"demo/util/models/token"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
)

func GetToken(uid uint64) string {
	ca := cache.NewCache()
	var model token.TokenModel
	model.ID = uid
	newID, _ := id.NewID(1)
	model.Token = newID.String()
	ca.Set(model.Token, cache.CacheModel{
		Time: cache.CacheDay,
		Data: model,
	})
	return model.Token
}
