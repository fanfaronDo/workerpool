package main

import (
	"sync"

	"github.com/fanfaronDo/workerpool/pkg/worker"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		w := worker.Worker{}
		w.Id = i
		go w.Process("hell", wg)
	}
	wg.Wait()
}
