package main

import (
	"time"
	"fmt"
)

func main() {
	c := make(chan int)
	c2 := make(chan int)

	go func() {
		c <- 1
	}()

	go func() {
		for {
			c <- <- c2
		}
	}()

	debug := time.NewTicker(time.Duration(time.Millisecond * 50))
	for {
		select {
		case v:= <-c:
			c2 <- v
		case <-debug.C:
			fmt.Println("debug ticker")
		}
	}
}
