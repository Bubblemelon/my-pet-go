package database

import "github.com/gomodule/redigo/redis"

// RedisConnection connects to the Redis server
// and returns the connect of type redis.Conn
func RedisConnection() redis.Conn {

	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}

	return c

}
