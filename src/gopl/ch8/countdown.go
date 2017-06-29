package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan bool)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- true
	}()

	fmt.Println("Press Enter to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown >= 0; countdown-- {
		fmt.Println(time.Now(), countdown)
		select {
		case <-tick:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted.")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
