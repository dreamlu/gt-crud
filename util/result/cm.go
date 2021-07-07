package result

import (
	"errors"
	"github.com/dreamlu/gt/tool/log"
	te "github.com/dreamlu/gt/tool/type/errors"
)

func Res(err error) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = MapSuccess
	}
	return
}

func ResGet(data interface{}, err error) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		res = GetSuccess(data)
	}
	return
}

func ResPager(datas interface{}, pager Pager, err error) (res Resultable) {
	if err != nil {
		res = CError(err)
		if !errors.As(err, &te.TextErr) {
			log.Error(res)
		}
	} else {
		if pager.TotalNum == 0 {
			datas = []interface{}{}
		}
		res = GetSuccessPager(datas, pager)
	}
	return
}
