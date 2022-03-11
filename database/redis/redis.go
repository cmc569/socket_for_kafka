package redis

import (
	"api/config/setting"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"time"
)

type connectRedis struct {
	connectRedis *redis.Pool
}

//Global Config Variable
var ConnectRedis connectRedis

// OpenConnect 返回redis连接池
func OpenConnect() {
	ConnectRedis.connectRedis = &redis.Pool{
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://" + setting.RedisConfig.Host + ":" + setting.RedisConfig.Port)

			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", setting.RedisConfig.Password); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}

			if _, dbErr := c.Do("SELECT", setting.RedisConfig.DBName); dbErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", dbErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

////////////////////////////
///       key-value      ///
////////////////////////////

func Set(key string, value interface{}) error {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	_, err := c.Do("SET", key, value)
	return err
}

func Get(key string) (string, error) {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	value, err := redis.String(c.Do("GET", key))
	return value, err
}

func DelKey(key string) error {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	_, err := c.Do("DEL", key)
	return err
}

func IsExistByKey(key string) bool {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

////////////////////////////
///       SETNX	         ///
////////////////////////////

func SetSetnx(key string, value interface{}) error {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	_, err := c.Do("SETNX", key, value)
	return err
}

func GetSetnx(key string) (string, error) {
	c := ConnectRedis.connectRedis.Get()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	return string(jsonGet), err
}

////////////////////////////
///       hex-table      ///
////////////////////////////

func HSet(table, key string, value interface{}) {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	_, err := c.Do("hset", table, key, value)
	if err != nil {
		fmt.Println("haset failed", err.Error())
	}
}

func HGet(table, key string) string {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	res, err := c.Do("hget", table, key)
	if err != nil {
		fmt.Println("hget failed", err.Error())
	}
	if res == nil {
		return ""
	} else {
		return string(res.([]byte))
	}
}

//var s struct {
//	Title  string `redis:"title"`
//	Author string `redis:"author"`
//	Body   string `redis:"body"`
//}
func HSetModel(s interface{}) {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	if _, err := c.Do("hmset", redis.Args{}.Add(reflect.TypeOf(s).Elem().Name()).AddFlat(s)...); err != nil {
		fmt.Println("HSetModel failed", err.Error())
	}
}

//獲取多個value
func HMGet(table string, keyAry []string) (map[string]string, error) {
	resMap := make(map[string]string, 0)
	var resErr error
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()

	for _, key := range keyAry {
		res, err := c.Do("hget", table, key)
		if err != nil {
			resErr = err
			return resMap, resErr
		} else if res == nil {
			return resMap, errors.New("res is nil")
		} else {
			resMap[key] = string(res.([]byte))
		}
	}
	return resMap, resErr
}

func HDelKey(table, key string) error {
	var err error
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	if _, err = c.Do("hdel", table, key); err != nil {
		fmt.Println("HSetModel failed", err.Error())
	}
	return err
}

func IsExistByHexTableKey(table, key string) bool {
	c := ConnectRedis.connectRedis.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("hexists", table, key))
	if err != nil {
		return false
	} else {
		return exist
	}
}


func GetHLen(table string) (int, error) {
	c := ConnectRedis.connectRedis.Get()
	dbHlen, err := c.Do("HLEN", table)

	return int(dbHlen.(int64)), err
}

func GetIncr(key string) int {
	c := ConnectRedis.connectRedis.Get()
	v, _ := redis.Int64(c.Do("INCR", key))
	return int(v)
}

func IncrInit(key string) {
	c := ConnectRedis.connectRedis.Get()
	c.Do("EXPIRE", key, 0)
}
