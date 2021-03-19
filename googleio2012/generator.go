package main

import (
	"fmt"
	"time"
)

// Generator: function that return a channel
func generator(msg string) <-chan string { // return read-only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(5 * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := generator("hello")
	g := generator("gawr gura")

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
		fmt.Println(<-g)
	}
	fmt.Println("Finished")
}
