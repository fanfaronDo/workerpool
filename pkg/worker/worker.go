package worker

import (
	"fmt"
	"log/slog"
	"time"
)

type Worker struct {
	Id   int
	Busy bool
}

func (w *Worker) Process(taskName string) {

	time.Sleep(1 * time.Second)
	slog.Info(fmt.Sprintf("Worker id is %d proccessed task: %s\n", w.Id, taskName))
}
