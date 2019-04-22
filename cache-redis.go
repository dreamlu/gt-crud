// @author  dreamlu
package deercoder

import (
	"bytes"
	"encoding/json"
	"strconv"
)

var (
	Host        = GetDevModeConfig("redis.host")
	Password    = GetDevModeConfig("redis.password")
	Database    = GetDevModeConfig("redis.database")
	MaxOpenConn = GetDevModeConfig("redis.maxOpenConn") // max number of connections
	MaxIdleConn = GetDevModeConfig("redis.maxIdleConn") // 最大的空闲连接数

	// redis conn
	// init it
	Rc *ConnPool
)

// in test, init no use
func init() {
	dba, _ := strconv.Atoi(Database)
	mops, _ := strconv.Atoi(MaxOpenConn)
	midas, _ := strconv.Atoi(MaxIdleConn)
	Rc = InitRedisPool(Host, Password, dba, mops, midas)
}

// impl cache manager
// redis cache
// interface key, interface value
type redisManager struct {
	// do nothing else
}

func (r *redisManager) set(key interface{}, value CacheModel) error {

	// change key to string
	keyS, err := json.Marshal(key)
	if err != nil {
		return err
	}

	// can not store struct data
	// change data to string
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// set string data
	err = Rc.Set(keyS, data).Err()
	if err != nil {
		return err
	}
	if value.Time != 0 {
		err = Rc.ExpireKey(keyS, value.Time*60).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *redisManager) get(key interface{}) (CacheModel, error) {

	var reply CacheModel

	// change key to string
	keyS, err := json.Marshal(key)
	if err != nil {
		return reply, err
	}

	// data
	res := Rc.Get(keyS).Val()

	// string to struct data
	err = json.Unmarshal([]byte(res.(string)), &reply)
	if err != nil {
		return reply, err
	}

	return reply, nil
}

func (r *redisManager) delete(key interface{}) error {

	// change key to string
	keyS, err := json.Marshal(key)
	if err != nil {
		return err
	}

	return Rc.Delete(keyS).Err()
}

func (r *redisManager) deleteMore(key interface{}) error {

	// change key to string
	keyS, err := json.Marshal(key)
	if err != nil {
		return err
	}

	var (
		buf bytes.Buffer
	)
	buf.WriteString("*")
	buf.Write(keyS)
	buf.WriteString("*")

	// keys
	res := Rc.Keys(buf.Bytes()).Val()
	if res != nil {
		for _,v := range res.([]interface{}) {
			err := Rc.Delete(v).Err()
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (r *redisManager) check(key interface{}) error {

	var reply CacheModel

	// change key to string
	keyS, err := json.Marshal(key)
	if err != nil {
		return err
	}

	// data
	res := Rc.Get(keyS).Val()

	// string to struct data
	err = json.Unmarshal([]byte(res.(string)), &reply)
	if err != nil {
		return err
	}

	return Rc.ExpireKey(keyS, reply.Time*60).Err()
}
