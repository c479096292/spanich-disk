package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	client   *redis.Client
	redisHost = "127.0.0.1:6379"
	redisPass = "testupload"
)

func newRedisPool()  {
	opt := redis.Options{
		Addr:redisHost,
		PoolSize: 20,
		DialTimeout: 5 * time.Second,
	}
	client = redis.NewClient(&opt)
	cmd := client.Do("ping")
	_, err := cmd.Result()
	if err !=nil {
		fmt.Println("connect redis error: ", err)
	}
}

func init() {
	newRedisPool()
}

//func main() {
//	stmt := client.Do("hgetall","person")
//	res, err := stmt.Result()
//	if err !=nil {
//		fmt.Println("err: ",err)
//	}
//	fmt.Printf("%T",res)
//	data, ok := res.([]interface{})
//	if !ok {
//		fmt.Println("error")
//	}
//	fmt.Println("res: ", data)
//}
func RedisPool() *redis.Client {
	return client
}