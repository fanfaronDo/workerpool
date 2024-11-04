package worker

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type Worker struct {
	Id int
}

func (w *Worker) Process(taskName string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	slog.Info(fmt.Sprintf("Worker id is %d proccessed task: %s\n", w.Id, taskName))
}
