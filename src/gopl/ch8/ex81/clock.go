package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 3 || args[1] != "-p" {
		fmt.Println("Usage:\n\tclock -p port\n\nGet time zone from environment variable TZ, e.g. TZ=Asia/Shanghai")
		return
	}

	// Get time zone from environment variable TZ
	fmt.Printf("TZ=%s\n", os.Getenv("TZ"))
	loc, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal(err)
		return
	}

	listener, err := net.Listen("tcp", "localhost:"+args[2])
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, loc) // handle one connection at a time
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	fmt.Println("Connected...")
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05"))
		// _, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05.000"))
		if err != nil {
			log.Print(err)
			return // e.g., client disconnected
		}
		time.Sleep(100 * time.Millisecond)
	}
}
