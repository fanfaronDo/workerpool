package workerpool

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/fanfaronDo/workerpool/pkg/worker"
)

type WorkerPool struct {
	workers        []*worker.Worker
	tasks          chan string
	addWorkerCh    chan *worker.Worker
	removeWorkerCh chan *worker.Worker
	wg             sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		tasks:          make(chan string),
		addWorkerCh:    make(chan *worker.Worker),
		removeWorkerCh: make(chan *worker.Worker),
	}
}

func (p *WorkerPool) Start() {
	go p.run()
}

func (p *WorkerPool) run() {
	for {
		select {
		case task := <-p.tasks:
			p.wg.Add(1)
			if len(p.workers) > 0 {
				worker := p.workers[0]
				go worker.Process(task, &p.wg)
			} else {
				slog.Info("No available workers to process the task.")
			}
		case newWorker := <-p.addWorkerCh:
			p.workers = append(p.workers, newWorker)
			slog.Info(fmt.Sprintf("Added Worker %d\n", newWorker.Id))
		case removeWorker := <-p.removeWorkerCh:
			for i, worker := range p.workers {
				if worker.Id == removeWorker.Id {
					p.workers = append(p.workers[:i], p.workers[i+1:]...)
					slog.Info(fmt.Sprintf("Removed Worker %d\n", removeWorker.Id))
					break
				}
			}
		}
	}
}

func (p *WorkerPool) AddWorker(worker *worker.Worker) {
	p.addWorkerCh <- worker
}

func (p *WorkerPool) RemoveWorker(worker *worker.Worker) {
	p.removeWorkerCh <- worker
}

func (p *WorkerPool) Submit(task string) {
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}
