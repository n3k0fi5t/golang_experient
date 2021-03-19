package main

func main() {
	c := make(chan int)

	// create a infinite loop do nothing
	go func () {
		for {

		}
	}()

	<- c
}
