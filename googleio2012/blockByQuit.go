package main

import "fmt"

func generator_with_quit(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				break
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := generator_with_quit("nek0fi5t", quit)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- true
}
