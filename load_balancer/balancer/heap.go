package balancer

// heap implementation
func (p WorkerPool) Len() int {
	return len(p)
}

func (p WorkerPool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *WorkerPool) Push(x interface{}) {
	n := len(*p)
	w := x.(*Worker)
	w.index = n
	*p = append(*p, w)
}

func (p *WorkerPool) Pop() interface{} {
	old := *p
	n := len(old)

	w := old[n-1]
	w.index = -1

	*p = old[:n-1]
	return w
}

func (p *WorkerPool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].index = i
	a[j].index = j
}
