package balancer

import (
	"container/heap"
	"fmt"
	"math"
	"time"
)

const (
	nRequest = 10
)

func InitBalancer(maxWorker int) *Balancer {
	// at least one worker
	maxWorker = int(math.Max(1.0, float64(maxWorker)))

	b := &Balancer{
		pool: make([]*Worker, 0, maxWorker),
		done: make(chan *Worker),
		pend: make(chan Request),
	}

	// create worker
	for i := 0; i < maxWorker; i++ {
		w := &Worker{
			work_queue: make(chan Request, nRequest),
		}
		// put worker in heap
		heap.Push(&b.pool, w)
		go w.doWark(b.done)
	}

	return b
}

func (b *Balancer) print() {
	for _, worker := range b.pool {
		fmt.Printf("%d ", worker.pending)
	}
	fmt.Printf("\n")
}

func (b *Balancer) Balance(req chan Request) {
	debug := time.NewTicker(time.Duration(time.Millisecond * 50))
	for {
		select {
		case request := <-req:
			b.dispatch(request)
		case w := <-b.done:
			b.complete(w)
		case <-debug.C:
			b.print()
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	worker := heap.Pop(&b.pool).(*Worker)
	if worker.pending >= nRequest {
		// drop request by send request back directly
		req.Res <- Result{Err: ResultError{"Request fail, worker was overload"}}
	} else {
		// this may cause deadlock while wocker's request queue is full
		worker.work_queue <- req
		worker.pending++
	}
	heap.Push(&b.pool, worker)
}

func (b *Balancer) complete(worker *Worker) {
	heap.Remove(&b.pool, worker.index)
	worker.pending--
	heap.Push(&b.pool, worker)
}

func (w *Worker) doWark(done chan *Worker) {
	for req := range w.work_queue {
		result := req.Fn()
		response := Result{Res: result, Err: nil}
		req.Res <- response
		done <- w
	}
}
