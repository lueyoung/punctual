package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
	"math/rand"

	"github.com/go-redis/redis"
)

var i = flag.String("i", "127.0.0.1", "IP of Redis")
var p1 = flag.String("1", "6379", "Port of Redis")
var p2 = flag.String("2", "6380", "Port of Redis")
var from = flag.String("f", "newbie", "todo")
var to = flag.String("t", "cached", "todo")
var n = flag.Int("n", 10000, "numbers")
var list = flag.String("l", "newbie", "Key of Redis list")

func init() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	hosts := strings.Split(*i, ",")
	l := len(hosts)
	addrs := make([]string, 2*l, 2*l)
	for i := 0; i < l; i++ {
		addrs[2*i] = fmt.Sprintf("%v:%v", hosts[i], *p1)
		addrs[2*i+1] = fmt.Sprintf("%v:%v", hosts[i], *p2)
	}
	log.Printf("Redis cluster addrs: %v\n", addrs)
	o := redis.ClusterOptions{
		Addrs: addrs,
	}
	client := redis.NewClusterClient(&o)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now().Unix()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *n; i++ {
		timestamp = time.Now().Unix()
		key := fmt.Sprintf("%v_%v", timestamp, rand.Int())
		value := fmt.Sprintf("%v", rand.Int())
		err = client.Set(key, value, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
		err = client.LPush(*list, key).Err()
		if err != nil {
			err2 := client.Del(key).Err()
			if err2 != nil {
				log.Fatal(err2)
			}
			log.Fatal(err)
		}
		if (i%10000 == 0 && i != 0) || i == 1 {
			fmt.Printf("data points written: %v\n", i)
			fmt.Printf("(%v, %v)\n", key, value)
		}
	}
}
