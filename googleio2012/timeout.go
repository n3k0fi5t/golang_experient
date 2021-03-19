package main

import (
	"fmt"
	"time"
)

func multiplex_with_timeout(in <-chan string) <-chan string {
	c := make(chan string)
	timeout := time.After(1 * time.Second)
	go func() {
		for {
			select {
			case s := <-in:
				c <- s
			case <-timeout:
				c <- "Timeout"
				close(c)
				return
			}
		}
	}()

	return c
}

func generator(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func main() {
	msg := generator("Hi~")
	c := multiplex_with_timeout(msg)

	for m := range c {
		fmt.Println(m)
	}

}
