package main

import (
	"flag"
	"fmt"
	"os"
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
		*cip = os.Getenv("KUBECONFIG")
	}
}

func main() {
	fmt.Printf("%v\n", *cip)
}
