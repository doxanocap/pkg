package workerpool

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool[T ITask, R IResult] struct {
	wg  *sync.WaitGroup
	job func(context.Context, T) R

	workersQty int
	tasksCh    chan T
	resultsCh  chan R
}

func NewWorkerPool[T ITask, R IResult](
	workersQty int,
	job func(context.Context, T) R,
) *WorkerPool[T, R] {
	wp := &WorkerPool[T, R]{
		job:        job,
		wg:         &sync.WaitGroup{},
		workersQty: workersQty,
		resultsCh:  make(chan R),
		tasksCh:    make(chan T),
	}

	return wp
}

func (pool *WorkerPool[T, R]) Start(ctx context.Context) {
	for i := 0; i < pool.workersQty; i++ {
		worker := SpanWorker[T, R](pool.tasksCh, pool.resultsCh, pool.job, pool.wg)

		workerDoneCh := make(chan struct{})
		worker.Start(ctx, workerDoneCh)

		go func(i int) {
			<-workerDoneCh
			pool.Log("worker %d is asleep", i)
			close(workerDoneCh)
		}(i)
	}

	pool.Log("started %d workers", pool.workersQty)
}

func (pool *WorkerPool[T, R]) Stop() {
	close(pool.tasksCh)
	pool.wg.Wait()
	close(pool.resultsCh)
}

func (pool *WorkerPool[T, R]) AppendTasks(tasks []T) {
	go func() {
		for i := 0; i < len(tasks); i++ {
			pool.tasksCh <- tasks[i]
		}
	}()
}

func (pool *WorkerPool[T, R]) GetResults() <-chan R {
	return pool.resultsCh
}

func (pool *WorkerPool[T, R]) Log(msg string, a ...any) {
	log.Println("[WORKER_POOL] " + fmt.Sprintf(msg, a...))
}
