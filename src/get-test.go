package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	//"github.com/gocql/gocql"
)

var rip = flag.String("r", "127.0.0.1", "IP of Redis")
var cip = flag.String("c", "127.0.0.1", "IP of Cassandra")
var k = flag.String("k", "", "Key to get")
var list = flag.String("l", "newbie", "Redis list to use")
var rport = flag.String("1", "6379", "Port of Redis")
var cport = flag.Int("2", 9042, "Port of Cassandra")
var keyspace = flag.String("3", "punctual", "Keyspace of Cassandra")
var table = flag.String("4", "history", "Table of Cassandra")

var _ = flag.Parse()

var addr = fmt.Sprintf("%s:%s", *rip, *rport)

//log.Printf("Redis addr: %v\n", addr)
var o = redis.Options{
	Addr:     addr,
	Password: "",
	DB:       0,
}
var client = redis.NewClient(&o)

func main() {
	defer client.Close()
	// 1 Get key from Redis
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	v, err := client.Get(*k).Result()
	if err != nil {
		fmt.Println(err)
	}
	if err.Error() == "redis: nil" {
		fmt.Println("bingo")
	}
	fmt.Printf("get: (%v, %v)\n", *k, v)
}
