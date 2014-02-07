package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	start := flag.NewFlagSet("start", flag.ExitOnError)
	dbIp := start.String("db_ip", "127.0.0.1", "ip address for databse host")
	startHelp := start.Bool("help", false, "show this help message")

	stop := flag.NewFlagSet("stop", flag.ExitOnError)
	flushDb := stop.Bool("flush_db", false, "flush database")

	if len(os.Args) == 1 {
		showUsage()
		return
	}

	if os.Args[1] == "start" {
		start.Parse(os.Args[2:])
		if *startHelp {
			start.PrintDefaults()
			return
		}
		fmt.Printf("Database ip address: %v\n", *dbIp)
		return
	}

	if os.Args[1] == "stop" {
		stop.Parse(os.Args[2:])
		if *flushDb {
			fmt.Println("stopping and flushing db")
			return
		}
		fmt.Println("stopping")
		return
	}
	showUsage()
}

func showUsage() {
	fmt.Println("Usage: {start|stop}")
}
