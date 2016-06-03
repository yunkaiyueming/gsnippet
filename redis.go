package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var rs redis.Conn

func init() {
	GetRedisConn()
}

func main() {
	//GetSetKey()
	//LpushGet()
	//LpushScan()
	HgetSet()
}

func GetRedisConn() {
	// tcp连接
	var err error
	dialoption := redis.DialPassword("admin")
	rs, err = redis.Dial("tcp", "localhost:6379", dialoption)

	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
	} else {
		fmt.Println("redis conn success")
	}
}

func CloseRedisConn() {
	rs.Close()
}

func GetSetKey() {
	// 选择db
	rs.Do("SELECT", 0)

	rs.Do("set", "goredis", "ok111")
	v, err := redis.String(rs.Do("get", "goredis"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

func LpushGet() {
	rs.Do("lpush", "redlist", "qqq")
	rs.Do("lpush", "redlist", "www")
	rs.Do("lpush", "redlist", "eee")

	values, _ := redis.Values(rs.Do("lrange", "redlist", "0", "100"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}

func LpushScan() {
	var v1 string
	values, _ := redis.Values(rs.Do("lrange", "redlist", "0", "100"))
	redis.Scan(values, &v1)
	fmt.Println(v1)
}

func HgetSet() {
	n, err := rs.Do("HSET", "QahTest", "name", "QiAihui")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	s, err := redis.String(rs.Do("HGET", "QahTest", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
