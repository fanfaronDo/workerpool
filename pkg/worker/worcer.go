package worker

import (
	"log/slog"
	"sync"
	"time"
)

type identificator int

type Worker struct {
	Identificator identificator
}

func (w *Worker) Process(taskName string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	slog.Info("Worker id is %t proccessed task: %s\n", w.Identificator, taskName)
}
