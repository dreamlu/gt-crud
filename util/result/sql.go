package result

import (
	"fmt"
	"github.com/dreamlu/gt/tool/result"
	errors2 "github.com/dreamlu/gt/tool/type/errors"
	"strings"
)

// sql CN prompt
func GetSQLError(er error) (err error) {

	erS := er.Error()
	switch {
	case erS == "record not found":
		err = fmt.Errorf("%w", &errors2.TextError{Msg: result.MsgNoResult})
	case strings.Contains(erS, "PRIMARY"):
		err = fmt.Errorf("%w", &errors2.TextError{Msg: "主键重复"})
	case strings.Contains(erS, "Duplicate entry"):
		errs := strings.Split(erS, "for key ")
		erS = strings.Trim(errs[1], "'")
		if strings.Contains(erS, ".") {
			erS = strings.Split(erS, ".")[1]
		}
		err = fmt.Errorf("%w", &errors2.TextError{Msg: erS})
	case strings.Contains(erS, "Error 1406") || strings.Contains(erS, "Error 1264"):
		key := strings.Split(strings.Split(erS, "column '")[1], "'")[0]
		err = fmt.Errorf("%w", &errors2.TextError{Msg: fmt.Sprintf("字段过长[%s]", key)})
	default:
		err = er
	}

	return err
}
