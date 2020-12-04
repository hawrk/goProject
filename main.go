package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goredis/common"
	"goredis/curd"
	"log"
	"os"
	"time"
)

func init() {
	// check log dir exist
	if _, err := os.Stat(common.LogPath); os.IsNotExist(err) {
		_ = os.Mkdir(common.LogPath, 0777)
		_ = os.Chmod(common.LogPath, 0777)
	}
	file := "./" + common.LogPath + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	common.Loger = log.New(logFile, "", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime)
}

func main() {

	nowTime := time.Now().UnixNano() / 1e6
	fmt.Println("nowTime:", nowTime)

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Errorf("connect  redis :127.0.0.1:6379 fail")
		common.Loger.Println("connect  redis :127.0.0.1:6379 fail|err:", err)
		return
	}
	defer c.Close()

	fmt.Println("start to set key")

	err = curd.SetRedisKeyValue(c, "hawrk", "hawrkchen@tencent.com")
	if err != nil {
		return
	}
	err = curd.SetRedisKeyValueWithTimeOut(c, "hawrk2012", "hawrk2012@163.com", 15)
	if err != nil {
		return
	}
}
