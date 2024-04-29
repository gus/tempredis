package tempredis

import (
	"github.com/gomodule/redigo/redis"
)

func ExampleStart() {
	server, err := Start(Config{"databases": "8"})
	if err != nil {
		panic(err)
	}
	defer server.Term()

	conn, err := redis.Dial("unix", server.Socket())
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	conn.Do("SET", "foo", "bar")
}
