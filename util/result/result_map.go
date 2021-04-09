package result

import (
	"encoding/json"
)

// Resultable interface
type Resultable interface {
	Add(key string, value interface{}) (rmp Resultable) // Add
	AddStruct(value interface{}) (rmp Resultable)       // AddStruct
	String() string                                     // String()
}

type ResultMap map[string]interface{}

func (c ResultMap) Add(key string, value interface{}) Resultable {
	c[key] = value
	return c
}

func (c ResultMap) AddStruct(value interface{}) Resultable {
	b, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil
	}
	return c
}

// impl String()
func (c ResultMap) String() string {
	return StructToString(c)
}

func NewResultMap() ResultMap {
	return make(ResultMap)
}
