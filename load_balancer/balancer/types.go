package balancer

type CustomerResult interface{}

type ResultError struct {
	reason string
}

func (e ResultError) Error() string {
	return string(e.reason)
}

type Result struct {
	Res CustomerResult
	Err error
}

type Request struct {
	Fn  func() CustomerResult
	Res chan Result
}

type Worker struct {
	work_queue chan Request
	pending    int //number of pedning jobs
	index      int //index of heap
}

type WorkerPool []*Worker

type Balancer struct {
	pool WorkerPool
	done chan *Worker // channel to receive "done" signal from worker
	pend chan Request // pending request channel
}
