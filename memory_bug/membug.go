package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 2) // buffered channel
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Counter value %d\n", count)
}
