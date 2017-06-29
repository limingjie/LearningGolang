package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("Connection from", conn.RemoteAddr(), "starts...")
	input := bufio.NewScanner(conn)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go func(shout string) {
			fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
			time.Sleep(1 * time.Second)
			fmt.Fprintln(conn, "\t", shout)
			time.Sleep(1 * time.Second)
			fmt.Fprintln(conn, "\t", strings.ToLower(shout))
			time.Sleep(1 * time.Second)
			wg.Done()
		}(input.Text())
	}

	wg.Wait()
	conn.Close()
	fmt.Println("Connection from", conn.RemoteAddr(), "ends.")
}
