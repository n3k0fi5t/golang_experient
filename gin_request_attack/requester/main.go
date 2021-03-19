package main

import "net/http"

const nRequest = 100000

func main() {
	pending := make(chan int)

	// create request
	for i := 0; i < nRequest; i++ {
		go func() {
			for {
				http.Get("http://localhost:8080/")
			}
		}()
	}

	<-pending
}
