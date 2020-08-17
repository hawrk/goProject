package curd

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/cast"
	"goredis/common"

)

func SetRedisKeyValueWithTimeOut(con redis.Conn, key string, value string , ttl int32) error {
	_, err := con.Do("SET", key, value, "EX", cast.ToString(ttl))
	if err != nil {
		fmt.Errorf("SET KEY Error")
		common.Loger.Println("SET KEY Error:", err)
		return err
	}
	fmt.Println("set Key :", key, " with time out success")
	common.Loger.Println("set Key :", key, " with time out success")
	return nil
}

func SetRedisKeyValue(con redis.Conn, key ,value string ) error {
	_, err := con.Do("SET", key, value)
	if err != nil {
		fmt.Errorf("SET KEY Error")
		common.Loger.Println("SET KEY Error:", err)
		return err
	}
	fmt.Println("set Key :", key, " success")
	common.Loger.Println("set Key :", key, " success")
	return nil
}
