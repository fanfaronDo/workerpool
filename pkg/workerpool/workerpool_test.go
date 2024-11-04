package workerpool_test

import (
	"fmt"
	"testing"

	"github.com/fanfaronDo/workerpool/pkg/worker"
	"github.com/fanfaronDo/workerpool/pkg/workerpool"
)

func TestWorkerPool(t *testing.T) {
	pool := workerpool.NewWorkerPool()
	pool.Start()

	for i := 1; i <= 3; i++ {
		pool.AddWorker(&worker.Worker{Id: i})
	}

	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _, task := range tasks {
		pool.Submit(task)
	}

	pool.Wait()

	fmt.Println(pool.GetWorkers())
	if len(pool.GetWorkers()) != 3 {
		t.Errorf("Expected 3 workers, got %d", len(pool.GetWorkers()))
	}
	pool.RemoveWorker(&worker.Worker{Id: 1})
	fmt.Println(pool.GetWorkers())
	if len(pool.GetWorkers()) != 2 {
		t.Errorf("Expected 2 workers after removal, got %d", len(pool.GetWorkers()))
	}

	pool.Submit("Task 4")
	pool.Submit("Task 5")

	pool.Wait()
}

func TestWorkerPoolNoWorkers(t *testing.T) {
	pool := workerpool.NewWorkerPool()
	pool.Start()
	pool.Submit("Task 1")

	pool.Wait()
	if len(pool.GetWorkers()) != 0 {
		t.Errorf("Expected 0 workers, got %d", len(pool.GetWorkers()))
	}
}