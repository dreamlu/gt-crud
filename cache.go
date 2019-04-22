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
	set(key interface{}, value CacheModel) error
	// get value
	get(key interface{}) (CacheModel, error)
	// delete value
	delete(key interface{}) error
	// more del
	// key will become *key*
	deleteMore(key interface{}) error
	// check value
	// flush the time
	check(key interface{}) error
}
