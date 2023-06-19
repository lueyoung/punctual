package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var ip = flag.String("i", "127.0.0.1", "IP of Rdis")
var port = flag.String("p", "6379", "Port of Rdis")
var list = flag.String("l", "newbie", "Key of the list")
var k = flag.String("k", "", "Key to put")
var v = flag.String("v", "", "Value to put")

func init() {
	flag.Parse()
}

func main() {
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	//log.Printf("Redis addr: %v\n", addr)
	o := redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&o)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Set(*k, *v, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	err = client.LPush(*list, *k).Err()
	if err != nil {
		err2 := client.Del(*k).Err()
		if err2 != nil {
			log.Fatal(err2)
		}
		log.Fatal(err)
	}
	fmt.Printf("written: (%v, %v)\n", *k, *v)
}
