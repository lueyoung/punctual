package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
)

var rip = flag.String("r", "127.0.0.1", "IP of Redis")
var rport1 = flag.String("2", "6379", "Port of Redis")
var rport2 = flag.String("3", "6380", "Port of Redis")

func init() {
	flag.Parse()
}

func main() {
	hosts := strings.Split(*rip, ",")
	l := len(hosts)
	addrs := make([]string, 2*l, 2*l)
	for i := 0; i < l; i++ {
		addrs[2*i] = fmt.Sprintf("%v:%v", hosts[i], *rport1)
		addrs[2*i+1] = fmt.Sprintf("%v:%v", hosts[i], *rport2)
	}
	//log.Printf("Redis cluster addrs: %v\n", addrs)
	o := redis.ClusterOptions{
		Addrs: addrs,
	}
	client := redis.NewClusterClient(&o)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	ret := client.ClusterNodes()
	if ret.Err() != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster nodes: %v\n", ret.Val())
}
