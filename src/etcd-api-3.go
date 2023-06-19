package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
	//"enconding/json"
	"go.etcd.io/etcd/clientv3"
)

var ip = flag.String("i", "10.254.247.1", "endpoint IP")
var port = flag.String("p", "2379", "endpoint port")

func main() {
	flag.Parse()
	log.Printf("endpoint ip: %s\n", *ip)
	log.Printf("endpoint port: %s\n", *port)
	endpoint := fmt.Sprintf("%v:%v", *ip, *port)
	log.Printf("endpoint: %s\n", endpoint)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	})
	log.Printf("err: %v\n", err)
	if err != nil {
		log.Fatalf("connect failed. err: %v\n", err)
		return
	}
	log.Println("connect succ")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, "test")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		log.Printf("%s: %s\n", ev.Key, ev.Value)
	}
}
