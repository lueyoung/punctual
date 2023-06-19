package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gocql/gocql"
)

var rip = flag.String("r", "127.0.0.1", "IP of Redis")
var cip = flag.String("c", "127.0.0.1", "IP of Cassandra")
var cport = flag.Int("1", 9042, "Port of Cassandra")
var rport1 = flag.String("2", "6379", "Port of Redis")
var rport2 = flag.String("3", "6380", "Port of Redis")
var n = flag.Int64("n", 0, "numbers")
var from = flag.String("f", "newbie", "todo")
var to = flag.String("t", "cached", "todo")
var sleep = flag.Int("s", 3, "sleep time")
var keyspace = flag.String("k", "punctual", "Keyspace of Cassandra")
var table = flag.String("a", "history", "Table of Cassandra")

func init() {
	flag.Parse()
	if *cip == "127.0.0.1" || *cip == "localhost" {
		log.Fatalf("dont use \"%v\" as the IP addr of Cassandra", *cip)
	}
}

func main() {
	hosts := strings.Split(*rip, ",")
	l := len(hosts)
	addrs := make([]string, 2*l, 2*l)
	for i := 0; i < l; i++ {
		addrs[2*i] = fmt.Sprintf("%v:%v", hosts[i], *rport1)
		addrs[2*i+1] = fmt.Sprintf("%v:%v", hosts[i], *rport2)
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
	var tmp int64
	var llen int64
	var pllen *int64
	for i := 0; i < 3; i++ {
		ret := client.LLen(*from)
		if ret.Err() != nil {
			log.Fatal("err", ret.Err())
			return
		}
		log.Printf("%v len: %v\n", *from, ret.Val())
		llen = ret.Val()
		if llen == 0 {
			log.Printf("%v len: %v, exit\n", *from, ret.Val())
			return
		}
		if i == 0 {
			tmp = llen
		} else {
			if tmp != llen {
				log.Printf("previous len: %v, current len: %v, exit\n", tmp, llen)
				return
			}
			log.Printf("previous len == current len: %v\n", llen)
		}
		pllen = &llen
		if i < 2 {
			time.Sleep(time.Duration(*sleep) * time.Second)
		}
	}
	log.Printf("copy data from Redis to Cassandra: %v\n", *pllen)
	ips := strings.Split(*cip, ",")
	//cluster := gocql.NewCluster(fmt.Sprintf("%v", *cip))
	log.Printf("Cassndra addr: %v\n", ips)
	cluster := gocql.NewCluster()
	cluster.Hosts = ips
	cluster.Keyspace = *keyspace
	cluster.Port = *cport
	log.Printf("Cassndra using: %v.%v\n", *keyspace, *table)
	//cluster.Consistency = gocql.Quorum
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for {
		i++
		cassandra := 0
		// 1 RPop from newbie, and get the key
		key, err := client.RPop(*from).Result()
		if err != nil {
			log.Fatal(err)
		}
		// 2 Get value from Redis with the key
		value, err := client.Get(key).Result()
		if err != nil {
			log.Println(err)
			log.Printf("Key %v in list %v, not in Redis, discard\n", key, *from)
			cassandra = 1
		}
		if cassandra == 0 {
			// 3 Write key, value to Cassandra
			timeline := gocql.TimeUUID()
			err2 := session.Query(`INSERT INTO history (key, value, timeline) VALUES (? , ?, ?)`, key, value, timeline).Exec()
			if err2 != nil {
				// 4.1 faild, RPush the kay to newbie
				err = client.RPush(*from, key).Err()
				if err != nil {
					log.Fatal(err)
				}
				log.Fatal(err2)
			}
			// 4.2 succ, LPush the key to cached
			/*
				err = client.LPush(*to, key).Err()
				if err != nil {
					log.Fatal(err)
				}*/
			if (i%10000 == 0 && i != 0) || i == 1 {
				fmt.Printf("data cp: %v\n", i)
				fmt.Printf("(%v, %v, %v)\n", key, value, timeline)
			}
		}
		ret := client.LLen(*from)
		if ret.Err() != nil {
			if fmt.Sprintf("%v", ret.Err()) == "redis: nil" {
				llen = 0
			} else {
				log.Fatal("err", ret.Err())
			}
		} else {
			llen = ret.Val()
		}
		if llen == 0 {
			log.Printf("empty %v, after %v terms dealed, exit\n", *from, i)
			return
		}
	}
}
