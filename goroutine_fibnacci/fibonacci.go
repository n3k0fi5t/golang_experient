package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case v := <-quit:
			fmt.Printf("stop with value %d\n", v)
			return
		}
	}
}

const (
	N int = 10
)

func main() {
	quit := make(chan int)
	c := make(chan int)

	go func(c, quit chan int, n int) {
		for i := 0; i < n; i++ {
			fmt.Printf("%d %d\n", i, <-c)
		}
		quit <- -1
	}(c, quit, N)
	fibonacci(c, quit)
}
