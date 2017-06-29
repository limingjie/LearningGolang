package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:\n\tclock location1=server:port [location2...]")
		return
	}

	go connectClocks(args)

	fmt.Println("Press Enter to exit...")
	os.Stdin.Read(make([]byte, 1))
}

func connectClocks(args []string) {
	locations := make([]string, 0, len(args)-1)
	timestamps := make([]string, 0, len(args)-1)
	for i, arg := range args[1:] {
		arr := strings.Split(arg, "=")
		if len(arr) != 2 {
			fmt.Println("Invalid argument", arg)
			continue
		}
		locations = append(locations, arr[0])
		timestamps = append(timestamps, "--:--:--")
		go process(arr[0], arr[1], &timestamps[i])
	}

	for {
		fmt.Printf("\r\t")
		for i, loc := range locations {
			fmt.Printf("%10s %s\t", loc, timestamps[i])
		}
		fmt.Printf("\r")
		time.Sleep(100 * time.Millisecond)
	}
}

func process(loc string, server string, timestamp *string) {
	data := make([]byte, 512)
	for {
		conn, err := net.Dial("tcp", server)
		if err != nil {
			time.Sleep(1 * time.Second) // wait 1s before reconnecting
			continue
		}
		for {
			n, err := conn.Read(data)
			if err != nil {
				*timestamp = "--:--:--"
				break
			}
			*timestamp = string(data[:n])
		}
		conn.Close()
	}
}
