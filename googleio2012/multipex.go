package main

import (
	"fmt"
)

func multiplex(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-in1
		}
	}() // read from in1 and put to c
	go func() {
		for {
			c <- <-in2
		}
	}() // read from in1 and put to c

	return c
}

func multiplex_select(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()
	return c
}

func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func main() {
	msg1 := generator("hello")
	msg2 := generator("gopher")

	c := multiplex(msg1, msg2)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	msg3 := generator("gawr")
	msg4 := generator("gura")
	c2 := multiplex_select(msg3, msg4)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
	}
}
