package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

var meminfo = flag.String("m", "/proc/meminfo", "path to meminfo")
var threshold = flag.Int64("t", 50, "threshold of meminfo")
var result = flag.String("r", "/tmp/ret", "file to store result")
var n = flag.Int64("n", 1000000, "the number of msg to delete")
var ip = flag.String("i", "127.0.0.1", "IP of Rdis")
var port = flag.String("p", "6379", "Port of Rdis")
var list = flag.String("l", "cached", "Key of the list")

func chk() (int, int64) {
	file, err := os.OpenFile(*meminfo, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	var kb [5]int64
	pat := regexp.MustCompile(`\d+`)
	for i := 0; i < 5; i++ {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		match := pat.FindAllStringSubmatch(line, -1)
		kb[i], _ = strconv.ParseInt(match[0][0], 10, 64)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}
	total := kb[0]
	free := kb[1]
	//avaiable 	:= kb[2]
	buffer := kb[3]
	cache := kb[4]
	mem_use := total - free - buffer - cache
	mem_use_in_g := mem_use / 1024 / 1024
	total_in_g := total / 1024 / 1024
	mem_use_in_p := mem_use * 100 / total
	log.Printf("mem use: %v GB\n", mem_use_in_g)
	log.Printf("total mem: %v GB\n", total_in_g)
	log.Printf("mem use: %v %%\n", mem_use_in_p)
	if mem_use_in_p > *threshold {
		return 1, mem_use_in_p
	}
	/*
		ret_file, err := os.OpenFile(*result, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer ret_file.Close()
	*/
	return 0, mem_use_in_p
}

func init() {
	flag.Parse()
}

func main() {
	log.Printf("meminfo: %v\n", *meminfo)
	trigger, mem_use := chk()
	if trigger == 0 {
		log.Printf("mem use is %v %%, below threshold %v %%, exit\n", mem_use, *threshold)
		return
	}
	log.Printf("mem use is %v %%, above threshold %v %%, remove data\n", mem_use, *threshold)
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	log.Printf("Redis addr: %v\n", addr)
	o := redis.ClusterOptions{
		Addrs: []string{addr},
	}
	client := redis.NewClusterClient(&o)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	/*
		ret := client.LLen(*list)
		if ret.Err() != nil {
			log.Fatal("err", ret.Err())
			return
		}
		llen := ret.Val()
		log.Printf("%v len: %v\n", *list, llen)
		if llen == 0 {
			return
		}
		if llen <= *n {
			var i int64
			for i = 0; i < llen; i++ {
				key, err := client.RPop(*list).Result()
				if err != nil {
					err2 := client.RPush(*list, key).Err()
					if err2 != nil {
						log.Fatal(err2)
					}
					log.Fatal(err)
				}
				value, err := client.Get(key).Result()
				if err != nil {
					err3 := client.RPush(*list, key).Err()
					if err3 != nil {
						log.Fatal(err3)
					}
					log.Fatal(err)
				}
				log.Printf("key: %v, value: %v\n", key, value)
				err = client.Del(key).Err()
				if err != nil {
					err4 := client.RPush(*list, key).Err()
					if err4 != nil {
						log.Fatal(err3)
					}
					log.Fatal(err)
				}
			}
			return
		}*/
	for true {
		ret := client.LLen(*list)
		if ret.Err() != nil {
			log.Fatal("err", ret.Err())
		}
		llen := ret.Val()
		if llen == 0 {
			log.Printf("no data in %s\n", *list)
			return
		}
		var loops int64
		if llen <= *n {
			loops = llen
		} else {
			loops = *n
		}
		var i int64
		for i = 0; i < loops; i++ {
			key, err := client.RPop(*list).Result()
			if err != nil {
				err2 := client.RPush(*list, key).Err()
				if err2 != nil {
					log.Fatal(err2)
				}
				log.Fatal(err)
			}
			_, err = client.Get(key).Result()
			if err != nil {
				err3 := client.RPush(*list, key).Err()
				if err3 != nil {
					log.Fatal(err3)
				}
				log.Fatal(err)
			}
			err = client.Del(key).Err()
			if err != nil {
				err4 := client.RPush(*list, key).Err()
				if err4 != nil {
					log.Fatal(err4)
				}
				log.Fatal(err)
			}
		}
		if llen <= *n {
			log.Printf("no data in %s\n", *list)
			return
		}
		trigger2, mem_use := chk()
		if trigger2 == 0 {
			log.Printf("mem use is %v %%, below threshold %v %%, exit\n", mem_use, *threshold)
			return
		}
		log.Printf("mem use is %v %%, above threshold %v %%, remove data again\n", mem_use, *threshold)
	}
}
