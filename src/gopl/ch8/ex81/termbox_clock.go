package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func termboxPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func drawClock(x, y, w, h int, location string) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, ' ', termbox.ColorRed, termbox.ColorYellow)
		}
	}

	termboxPrint(x+1, y, termbox.ColorRed, termbox.ColorYellow, location)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:\n\tclock location1=server:port [location2...]")
		return
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	x, y := 2, 2
	for _, arg := range args[1:] {
		arr := strings.Split(arg, "=")
		if len(arr) != 2 {
			// fmt.Println("Invalid argument", arg)
			continue
		}
		drawClock(x, y, 12, 4, string(arr[0]))
		go client(arr[1], x+2, y+2)
		x += 16
	}

	const coldef = termbox.ColorDefault
	termboxPrint(0, 0, coldef, coldef, "Press ESC to exit...")
	go flush(100 * time.Millisecond)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			default:
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func flush(t time.Duration) {
	for {
		termbox.Flush()
		time.Sleep(t)
	}
}

func client(server string, x, y int) {
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
				termboxPrint(x, y, termbox.ColorRed, termbox.ColorYellow, "--:--:--")
				break
			}
			termboxPrint(x, y, termbox.ColorRed, termbox.ColorYellow, string(data[:n]))
		}
		conn.Close()
	}
}
