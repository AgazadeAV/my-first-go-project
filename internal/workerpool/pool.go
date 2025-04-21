package workerpool

import (
	"log"
	"sync"
)

type Pool struct {
	workers int
	jobs    chan Job
	wg      sync.WaitGroup
}

func NewPool(workers int) *Pool {
	return &Pool{
		workers: workers,
		jobs:    make(chan Job),
	}
}

func (pool *Pool) Start() {
	for i := 0; i < pool.workers; i++ {
		pool.wg.Add(1)
		go func(id int) {
			defer pool.wg.Done()
			for job := range pool.jobs {
				log.Printf("Worker %d processing job", id)
				job.CustomExecute()
			}
		}(i + 1)
	}
}

func (pool *Pool) Submit(job Job) {
	pool.jobs <- job
}

func (pool *Pool) Shutdown() {
	close(pool.jobs)
	pool.wg.Wait()
}
