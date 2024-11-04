package main

import (
	"fmt"

	"github.com/fanfaronDo/workerpool/pkg/worker"
	"github.com/fanfaronDo/workerpool/pkg/workerpool"
)

func main() {
	pool := workerpool.NewWorkerPool()
	pool.Start()

	for i := 1; i <= 5; i++ {
		pool.AddWorker(&worker.Worker{Id: i})
	}

	for i := 1; i <= 5; i++ {
		pool.Submit(fmt.Sprintf("Task %d", i))
	}

	pool.Wait()

	pool.RemoveWorker(&worker.Worker{Id: 2})

	for i := 6; i <= 10; i++ {
		pool.Submit(fmt.Sprintf("Task %d", i))
	}

	pool.Wait()
}
