// @author  dreamlu
package deercoder

// data model
type CacheModel struct {
	// minute
	Time int64 `json:"time"`
	// data
	Data interface{} `json:"data"`
}

// cache manager
type CacheManager interface {
	// operate method
	// set value
	// if time != 0 set it
	Set(key interface{}, value CacheModel) error
	// get value
	Get(key interface{}) (CacheModel, error)
	// delete value
	Delete(key interface{}) error
	// more del
	// key will become *key*
	DeleteMore(key interface{}) error
	// check value
	// flush the time
	Check(key interface{}) error
}
