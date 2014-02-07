package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.NewFlagSet("flag", flag.ExitOnError)
	dbIp := f.String("db_ip", "127.0.0.1", "ip address for databse host")
	help := f.Bool("help", false, "show this help message")
	f.Parse(os.Args[1:])
	if *help {
		f.PrintDefaults()
		return
	}
	fmt.Printf("Database ip address: %v\n", *dbIp)
}
