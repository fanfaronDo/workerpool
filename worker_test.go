package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id int
}

func (w *Worker) process(task string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d processed task: %s\n", w.id, task)
}

type WorkerPool struct {
	workers        []*Worker
	tasks          chan string
	addWorkerCh    chan *Worker
	removeWorkerCh chan *Worker
	wg             sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		tasks:          make(chan string),
		addWorkerCh:    make(chan *Worker),
		removeWorkerCh: make(chan *Worker),
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
				worker := p.workers[0] // Используем первого воркера
				go worker.process(task, &p.wg)
			} else {
				fmt.Println("No available workers to process the task.")
			}
		case newWorker := <-p.addWorkerCh:
			p.workers = append(p.workers, newWorker)
			fmt.Printf("Added Worker %d\n", newWorker.id)
		case removeWorker := <-p.removeWorkerCh:
			for i, worker := range p.workers {
				if worker.id == removeWorker.id {
					p.workers = append(p.workers[:i], p.workers[i+1:]...)
					fmt.Printf("Removed Worker %d\n", removeWorker.id)
					break
				}
			}
		}
	}
}

func (p *WorkerPool) AddWorker(worker *Worker) {
	p.addWorkerCh <- worker
}

func (p *WorkerPool) RemoveWorker(worker *Worker) {
	p.removeWorkerCh <- worker
}

func (p *WorkerPool) Submit(task string) {
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func main() {
	pool := NewWorkerPool()
	pool.Start()

	for i := 1; i <= 3; i++ {
		pool.AddWorker(&Worker{id: i})
	}

	for i := 1; i <= 5; i++ {
		pool.Submit(fmt.Sprintf("Task %d", i))
	}

	pool.Wait()

	pool.RemoveWorker(&Worker{id: 1})

	for i := 6; i <= 10; i++ {
		pool.Submit(fmt.Sprintf("Task %d", i))
	}

	pool.Wait()
}
