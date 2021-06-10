package token

import (
	"demo/util/models/token"
	"errors"
	"github.com/dreamlu/gt/cache"
	"github.com/dreamlu/gt/tool/id"
	"github.com/dreamlu/gt/tool/log"
)

var (
	ca = cache.NewCache()
)

func GetToken(model token.TokenModel) string {
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

// Expire 延长token
func Expire(token string) error {
	cam, err := ca.Get(token)
	if err != nil {
		return errors.New("token错误:" + err.Error())
	}
	if cam.Data == nil {
		return errors.New("token不存在,请重新登录获取")
	}
	// 延长token对应时间
	return ca.Set(token, cam)
}

func GetRole(tk string) ([]string, string, error) {

	model, err := GetTokenModel(tk)
	if err != nil {
		return nil, "", err
	}
	return model.Role, model.G, nil
}

func GetRoles(tk string) []string {

	model, err := GetTokenModel(tk)
	if err != nil {
		return []string{}
	}
	return model.Role
}

func GetTokenModel(tk string) (tm token.TokenModel, err error) {
	cam, err := ca.Get(tk)
	if err != nil {
		return tm, errors.New("token错误:" + err.Error())
	}
	err = cam.Unmarshal(&tm)
	if err != nil {
		return tm, errors.New("token错误:" + err.Error())
	}
	return
}
