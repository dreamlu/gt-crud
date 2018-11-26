package deercoder

import (
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	DefaultLogger            = Logger{log.New(os.Stdout, "\r\n", 0)}
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
	islog					 = 0
)

//设置文件log
func SetLogger(f *os.File) Logger{
	//return Logger{log.New(f, "DeerCoderSQL "+time.Now().Format("2006-01-02 15:04:05")+"\r\n", 0)}
	return Logger{log.New(f, "", 0)}
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {

		var (
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = "DeerCoderSQL TIME [" + time.Now().Format("2006-01-02 15:04:05") + "]\r\n"
			source          = fmt.Sprintf(" DeerCoderSQL FILE (%v)\r\n", values[1])
		)

		messages = []interface{}{source, currentTime}
		errInfo := fmt.Sprintf("%s",values[2])
		if level == "log" {
			if strings.Contains(errInfo,"Error") {
				islog = 1
			} else {
				//messages = nil
				return nil
			}

			messages = append(messages, "DeerCoderSQL SQL INFO")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\r\n")

		} else if level == "sql" && islog == 1 {
			islog = 0

			// duration
			messages = append(messages, fmt.Sprintf("DeerCoderSQL EXECUTE TIME [%.2fms] SQL", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf("\r\n DeerCoderSQL ROWS [%v]\r\n ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
		} else {
			return nil
		}
	}

	return
}

type logger interface {
	Print(v ...interface{})
}

// LogWriter log writer interface
type LogWriter interface {
	Println(v ...interface{})
}

// Logger default logger
type Logger struct {
	LogWriter
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	//format
	format := LogFormatter(values...)
	if nil == format {//过滤空白行
		return
	}
	logger.Println(format...)
}

