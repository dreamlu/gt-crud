package result

import (
	"github.com/dreamlu/gt/tool/type/bmap"
	"reflect"
)

// Resultable interface
type Resultable interface {
	Add(key string, v interface{}) (rmp Resultable) // Add
	AddStruct(v interface{}) (rmp Resultable)       // AddStruct
	String() string                                 // String()
}

type ResultMap bmap.BMap

func (c ResultMap) Add(key string, value interface{}) Resultable {
	c = ResultMap(bmap.BMap(c).Set(key, value))
	return c
}

func (c ResultMap) AddStruct(v interface{}) Resultable {
	if reflect.ValueOf(v).IsNil() {
		return c
	}
	switch v.(type) {
	case ResultMap:
		c.AddResultMap(v.(ResultMap))
	default:
		nr := bmap.StructToBMap(v)
		for k, v := range nr {
			c.Add(k, v)
		}
	}

	return c
}

func (c *ResultMap) AddResultMap(v ResultMap) {

	for k, v := range v {
		c.Add(k, v)
	}
	return
}

// impl String()
func (c ResultMap) String() string {
	return StructToString(c)
}

func NewResultMap() *ResultMap {
	return &ResultMap{}
}

func Add(key string, value interface{}) Resultable {
	return NewResultMap().Add(key, value)
}
