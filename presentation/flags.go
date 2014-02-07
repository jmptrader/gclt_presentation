package main

import (
	"flag"
	"fmt"
)

func main() {
	var dbIp = flag.String("db_ip", "127.0.0.1", "ip address for databse host")
	flag.Parse()
	fmt.Printf("Database ip address: %v\n", *dbIp)
}
