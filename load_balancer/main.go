package main

import (
	"math/rand"
	"time"
	"training_myself/load_balancer/balancer"
)

func createRequest() balancer.Request {
	resp := make(chan balancer.Result)
	work_func := func() balancer.CustomerResult {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return rand.Intn(10)
	}
	return balancer.Request{Fn: work_func, Res: resp}
}

func main() {
	req_chan := make(chan balancer.Request)
	b := balancer.InitBalancer(10)

	// create jobs randomly
	go func() {
		for {
			go func() {
				req := createRequest()
				// send request to balancer
				req_chan <- req
				// receiver request from worker
				<-req.Res
			}()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
	}()

	b.Balance(req_chan)
}
