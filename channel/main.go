package main

import (
	"fmt"
)

func block_by_writing_nil_channel() {
	var c chan int
	for {
		fmt.Println("block by nil writing channel ")
		c <- 1
	}
}

func block_by_read_nil_channel() {
	var c chan int
	for {
		fmt.Println("block by nil read channel ")
		<-c
	}
}

func write_to_closed_channel() {
	c := make(chan int)
	close(c) //close immediately
	for {
		fmt.Println("write to closed channel")
		c <- 1
	}
}

func read_from_closed_channel() {
	c := make(chan int)
	close(c) //close immediately
	for {
		fmt.Println("read from closed channel")
		t := <-c
	}
}

func main() {
	c := make(chan bool)

	go block_by_read_nil_channel()
	go block_by_writing_nil_channel()
	go read_from_closed_channel()
	//go write_to_closed_channel() //panic

	//block here
	<-c
}
