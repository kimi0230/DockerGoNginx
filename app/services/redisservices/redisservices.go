package redisservices

import (
	"fmt"
	"DockerGoNginx/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisClient : redis resouse pool
var RedisClient *redis.Pool

// MaxIdle：最大的空閒連接數,表示即使沒有redis連接時,依然可以保持N個空閒連接,而不會被清除,隨時待命狀態
// MaxActive：最大的連接數, 表示同時最多有N個連接
// IdleTimeout： 最大的空閒連接等待時間, 超過此時間,空閒連接會被關閉
func init() {
	RedisHost := config.REDIS_HOST + ":" + config.REDIS_PORT
	RedisDB := config.REDIS_DB
	RedisPassword := config.REDIS_PASSWORD

	// 建立連接
	RedisClient = &redis.Pool{
		MaxIdle:     config.MAX_IDLE,
		MaxActive:   config.MAX_ACTIVE,
		IdleTimeout: config.IDLE_TIMEOUT * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if RedisPassword != "" {
				if _, err := c.Do("AUTH", RedisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", RedisDB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},

		//定期對 redis server 做 ping/pong 測試
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

/*
Ping : ping pong
*/
func Ping() string {
	rc := RedisClient.Get()
	defer rc.Close()

	pong, err := rc.Do("PING")
	if err != nil {
		return ""
	}
	return pong.(string)
}

func demoHash() string {
	rc := RedisClient.Get()
	defer rc.Close()

	var myName string

	if _, err := rc.Do("hset", "me", "name", "kimi"); err != nil {
		fmt.Println("haset failed", err.Error())
	}
	res, err := rc.Do("hget", "me", "name")
	if err != nil {
		fmt.Println("hget failed", err.Error())
	} else {
		fmt.Printf("hget value :%s\n", res.([]byte))
	}
	rc.Do("expire", "me", "10")
	myName = string(res.([]byte))
	// rc.Do("hdel", "me", "name")
	return myName
}
