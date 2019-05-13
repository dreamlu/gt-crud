// @author  dreamlu
package deercoder

import (
	"log"
	"testing"
)

// redis method set test
func TestRedis(t *testing.T) {
	err := Rc.Set("test", "testValue").Err()
	log.Println("set err:", err)
	value := Rc.Get("test")
	reqRes,_ := value.Result()
	log.Println("value",reqRes)

}

// cache test
var cache CacheManager = new(RedisManager)

// user model
var user = User{
	ID:   1,
	Name: "test",
	//Createtime: JsonDate(time.Now()),
}

// set and get interface value
func TestCache(t *testing.T) {
	// data
	data := CacheModel{
		Time: 50,
		Data: user,
	}

	// key can use user.ID,user.Name,user
	// because it can be interface
	// set
	err := cache.Set(user, data)
	log.Println("set err: ", err)

	// get
	reply,_ := cache.Get(user)
	log.Println("user data :", reply.Data)

}

// check or delete cache
func  TestCacheCheckDel(t *testing.T)  {
	// check
	//err := cache.Check(user.ID)
	//log.Println("check: ", err)

	// del
	//err := cache.Delete(user.ID)
	//log.Println("delete: ", err)

	// del *

	//err := cache.Delete("1*")
	//log.Println("delete: ", err)

	// del more
	err := cache.DeleteMore(user)
	log.Println("delete: ", err)
}
