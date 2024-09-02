package workerpool

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool[T ITask] struct {
	wg  *sync.WaitGroup
	job func(context.Context, T) error

	workersQty int

	tasksCh   chan T
	resultsCh chan error

	maxRetries int
}

func NewWorkerPool[T ITask](
	workersQty int,
	job func(context.Context, T) error,
	options ...Option[T],
) *WorkerPool[T] {
	wp := &WorkerPool[T]{
		job:        job,
		wg:         &sync.WaitGroup{},
		workersQty: workersQty,
		tasksCh:    make(chan T),
		resultsCh:  make(chan error),
		maxRetries: 1,
	}

	for _, fn := range options {
		fn(wp)
	}

	return wp
}

func (pool *WorkerPool[T]) Start(ctx context.Context, tasks []T) {
	for i := 0; i < pool.workersQty; i++ {
		worker := SpanWorker[T](
			pool.tasksCh,
			pool.resultsCh,
			pool.job,
			pool.wg,
			pool.maxRetries)

		workerDoneCh := make(chan struct{})
		worker.Start(ctx, workerDoneCh)

		go func(i int) {
			<-workerDoneCh
			pool.Log("worker %d is asleep", i)
			close(workerDoneCh)
		}(i)
	}

	pool.Log("started %d workers", pool.workersQty)
	go func() {
		for i := 0; i < len(tasks); i++ {
			pool.tasksCh <- tasks[i]
		}
		pool.Stop()
	}()
}

func (pool *WorkerPool[T]) Stop() {
	close(pool.tasksCh)
	pool.wg.Wait()
	close(pool.resultsCh)
}

func (pool *WorkerPool[T]) GetResultsCh() <-chan error {
	return pool.resultsCh
}

func (pool *WorkerPool[T]) Log(msg string, a ...any) {
	log.Println("[WORKER_POOL] " + fmt.Sprintf(msg, a...))
}
