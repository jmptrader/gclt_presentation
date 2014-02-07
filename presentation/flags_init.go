package main

import (
	"flag"
	"fmt"
)

var dbIp = flag.String("db_ip", "127.0.0.1", "ip address for databse host")

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("Database ip address: %v\n", *dbIp)
}
