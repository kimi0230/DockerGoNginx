package config

import (
	"fmt"
	"os"
	"strconv"
)

/*
WRITE_LOG : 將API log 寫入DB
*/
const WRITE_LOG = true

/*
WRITE_LOG_TO_DB : 將API log 寫入DB
*/
const WRITE_LOG_TO_DB = true

/*
WRITE_FCM_LOG : 將FCM log 寫入DB
*/
const WRITE_FCM_LOG_TO_DB = true

/*
VALIDATION_VERSION :  GIN  Binding 驗證版本
	v8: app/validation/customValidateV9/customValidateV8.go
	v9: app/validation/customValidateV9/customValidateV9.go
*/
const VALIDATION_VERSION = "v9"

/* REDIS CONFIG */
var REDIS_HOST string = "127.0.0.1"
var REDIS_PORT string = "6379"
var REDIS_PASSWORD string = ""
var REDIS_DB int = 1

// MaxIdle：最大的空閒連接數,表示即使沒有redis連接時,依然可以保持N個空閒連接,而不會被清除,隨時待命狀態
const MAX_IDLE = 1

// MaxActive：最大的連接數, 表示同時最多有N個連接
const MAX_ACTIVE = 10

// IdleTimeout： 最大的空閒連接等待時間, 超過此時間,空閒連接會被關閉, Second
const IDLE_TIMEOUT = 180

func init() {
	if tmp := os.Getenv("REDIS_HOST"); tmp != "" {
		REDIS_HOST = tmp
	}
	if tmp := os.Getenv("REDIS_PORT"); tmp != "" {
		REDIS_PORT = tmp
	}

	if tmp := os.Getenv("REDIS_PASSWORD"); tmp != "" {
		REDIS_PASSWORD = tmp
	}
	fmt.Println("REDIS_PASSWORD", REDIS_PASSWORD)
	if tmp := os.Getenv("REDIS_DB"); tmp != "" {
		if tmpInt, err := strconv.Atoi(tmp); err != nil {
			REDIS_DB = tmpInt
		}
	}
}
