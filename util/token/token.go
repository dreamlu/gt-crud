package token

import (
	"demo/util/models/token"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/log"
)

func GetToken(uid uint64, typ int8) string {
	ca := cache.NewCache()
	var model token.TokenModel
	model.ID = uid
	model.Typ = typ
	newID, _ := id.NewID(1)
	model.Token = newID.String()
	err := ca.Set(model.Token, cache.CacheModel{
		Time: cache.CacheDay,
		Data: model,
	})
	if err != nil {
		log.Error(err)
	}
	return model.Token
}
