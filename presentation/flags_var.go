package main

import (
	"flag"
	"fmt"
)

func main() {
	var dbIp string
	flag.StringVar(&dbIp, "db_ip", "127.0.0.1", "ip address for database host")
	flag.Parse()
	fmt.Printf("Database ip address: %v\n", dbIp)
}
