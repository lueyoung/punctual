package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/gocql/gocql"
)

var rip = flag.String("r", "127.0.0.1", "IP of Redis")
var cip = flag.String("c", "127.0.0.1", "IP of Cassandra")
var k = flag.String("k", "", "Key to get")
var list = flag.String("l", "newbie", "Redis list to use")
var rport = flag.String("1", "6379", "Port of Redis")
var cport = flag.Int("2", 9042, "Port of Cassandra")
var keyspace = flag.String("3", "punctual", "Keyspace of Cassandra")
var table = flag.String("4", "history", "Table of Cassandra")
var ck = flag.String("5", "key", "Key index used in Cassandra")
var cv = flag.String("6", "value", "Value index used in Cassandra")

var reput chan int = make(chan int)

func init() {
	flag.Parse()
	if *cip == "127.0.0.1" || *cip == "localhost" {
		*cip = os.Getenv("POD_IP")
	}
}

func get_from_cassandra() string {
	cluster := gocql.NewCluster(fmt.Sprintf("%v", *cip))
	cluster.Keyspace = *keyspace
	cluster.Port = *cport
	//log.Printf("Cassndra using: %v.%v\n", *keyspace, *table)
	//cluster.Consistency = gocql.Quorum
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	defer session.Close()
	var v string
	err = session.Query(`SELECT value FROM punctual.history WHERE key = ? LIMIT 1`, *k).Consistency(gocql.One).Scan(&v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func reput_to_redis(client *redis.ClusterClient, v string) {
	err := client.Set(*k, v, 0).Err()
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
	//fmt.Printf("re-written: (%v, %v)\n", *k, v)
	reput <- 0
}

func main() {
	// 1 Get key from Redis
	addr := fmt.Sprintf("%s:%s", *rip, *rport)
	//log.Printf("Redis addr: %v\n", addr)
	o := redis.ClusterOptions{
		Addrs: []string{addr},
	}
	client := redis.NewClusterClient(&o)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	v, err := client.Get(*k).Result()
	if err == nil {
		// 2.1 if succ, ret
		fmt.Printf("%v,%v\n", *k, v)
		return
	}
	if err.Error() != "redis: nil" {
		// 2.2 if failed with other reasons, ret
		log.Fatal(err)
	}
	// 2.3 if not exists, get from Cassandra
	//log.Println("not in Redis, try Cassandra")
	v = get_from_cassandra()
	go reput_to_redis(client, v)
	<-reput
	fmt.Printf("%v,%v\n", *k, v)
}
