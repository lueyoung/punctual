package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

var ip = flag.String("i", "127.0.0.1", "endpoint IP")
var port = flag.String("p", "2379", "endpoint port")
var namespace = flag.String("n", "default", "namespace")
var t = flag.Int64("t", 300, "time to live")
var dirname = flag.String("d", "", "path name")
var pod_ip = flag.String("1", "", "pod ip")
var host_ip = flag.String("2", "", "host ip")
var host_hash = flag.String("3", "", "host hash")
var host_id = flag.String("4", "", "host id")

func main() {
	flag.Parse()
	ttl := time.Duration(*t) * time.Second
	log.Printf("endpoint ip: %s\n", *ip)
	log.Printf("endpoint port: %s\n", *port)
	endpoint := fmt.Sprintf("http://%v.%v:%v", *ip, *namespace, *port)
	log.Printf("endpoint: %s\n", endpoint)
	key := fmt.Sprintf("/%v/%v", *dirname, *host_hash)
	value := fmt.Sprintf("{\"podIP\":\"%v\",\"hostIP\":\"%v\",\"hostHash\":\"%v\",\"hostID\":\"%v\"}", *pod_ip, *host_ip, *host_hash, *host_id)
	log.Printf("key: %s, value: %v\n", key, value)

	cfg := client.Config{
		Endpoints:               []string{endpoint},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	o := client.SetOptions{
		TTL: ttl,
	}
	_, err = kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Printf("no %v found, create", key)
		_, err2 := kapi.Set(context.Background(), key, value, &o)
		if err2 != nil {
			log.Fatal(err)
		}
		return
	}
	log.Printf("%v found, refresh ttl", key)
	_, err = kapi.Set(context.Background(), key, value, &o)
	if err != nil {
		log.Fatal(err)
	}
}
