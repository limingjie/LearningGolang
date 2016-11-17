package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Println("Recovered:", str)
	}()
	panic("Panic!")
}
