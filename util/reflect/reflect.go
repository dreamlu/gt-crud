package reflect

import (
	"reflect"
)

// reflect new type value
func New(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return reflect.New(t).Interface()
}

// reflect new array type value
func NewArray(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	//log.Println(reflect.MakeSlice(reflect.SliceOf(t), 0, 0))
	return reflect.New(reflect.SliceOf(t)).Interface()
}

// reflect Annotate
//func Annotate(v interface{}) string {
//	t := reflect.TypeOf(v)
//	if t.Kind() == reflect.Ptr {
//		t = t.Elem()
//	}
//}
