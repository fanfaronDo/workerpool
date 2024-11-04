package workerpool

import (
	"fmt"
	"log/slog"
	"sync"

	myworker "github.com/fanfaronDo/workerpool/pkg/worker"
)

type WorkerPool struct {
	workers        []*myworker.Worker
	tasks          chan string
	addWorkerCh    chan *myworker.Worker
	removeWorkerCh chan *myworker.Worker
	wg             sync.WaitGroup
	mu             sync.Mutex
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		tasks:          make(chan string),
		addWorkerCh:    make(chan *myworker.Worker),
		removeWorkerCh: make(chan *myworker.Worker),
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
			go p.processTask(task)
		case newWorker := <-p.addWorkerCh:
			p.mu.Lock()
			p.workers = append(p.workers, newWorker)
			slog.Info(fmt.Sprintf("Added Worker %d\n", newWorker.Id))
			p.mu.Unlock()
		case removeWorker := <-p.removeWorkerCh:
			p.mu.Lock()
			for i, worker := range p.workers {
				if worker.Id == removeWorker.Id {
					p.workers = append(p.workers[:i], p.workers[i+1:]...)
					slog.Info(fmt.Sprintf("Removed Worker %d\n", removeWorker.Id))
					break
				}
			}
			p.mu.Unlock()
		}
	}
}

func (p *WorkerPool) processTask(task string) {
	defer p.wg.Done()
	worker := p.getFreeWorker()
	if worker != nil {
		worker.Busy = true
		worker.Process(task)
		worker.Busy = false
	} else {
		p.tasks <- task
	}
}

func (p *WorkerPool) getFreeWorker() *myworker.Worker {
	for _, worker := range p.workers {
		if !worker.Busy {
			return worker
		}
	}
	return nil
}

func (p *WorkerPool) AddWorker(worker *myworker.Worker) {
	p.addWorkerCh <- worker
}

func (p *WorkerPool) GetWorkers() []*myworker.Worker {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.workers
}

func (p *WorkerPool) RemoveWorker(worker *myworker.Worker) {
	p.removeWorkerCh <- worker
}

func (p *WorkerPool) Submit(task string) {
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}
