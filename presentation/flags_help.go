package main

import (
	"flag"
	"fmt"
)

func main() {
	var dbIp string
	var help bool
	flag.StringVar(&dbIp, "db_ip", "127.0.0.1", "ip address for database host")
	flag.BoolVar(&help, "help", false, "show this help message")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Database ip address: %v\n", dbIp)
}
