package workerpool

import (
	"context"
	"fmt"
	"log"
	"sync"
)

const ()

type Worker[T ITask] struct {
	tasksCh    <-chan T
	resultsCh  chan<- error
	job        func(context.Context, T) error
	wg         *sync.WaitGroup
	maxRetries int
}

func SpanWorker[T ITask](
	tasksCh <-chan T,
	resultsCh chan<- error,
	job func(ctx context.Context, task T) error,
	wg *sync.WaitGroup,
	maxRetries int,
) *Worker[T] {
	return &Worker[T]{
		job:        job,
		tasksCh:    tasksCh,
		resultsCh:  resultsCh,
		wg:         wg,
		maxRetries: maxRetries,
	}
}

func (w *Worker[T]) Start(ctx context.Context, doneCh chan<- struct{}) {
	go func(ctx context.Context, doneCh chan<- struct{}) {
		defer func() {
			doneCh <- struct{}{}
			w.wg.Done()
		}()

		w.wg.Add(1)
		for {
			select {
			case task, ok := <-w.tasksCh:
				if !ok {
					return
				}

				err := w.job(ctx, task)
				if err != nil {
					err = fmt.Errorf("task %d: %w", task.ID(), err)
				}

				w.resultsCh <- err

				// TODO refactor retries
				//succeed := false
				//sleepDur := 5 * time.Second
				//for i := 0; i < w.maxRetries; i++ {
				//	if err := w.job(ctx, task); err == nil {
				//		succeed = true
				//		break
				//	} else {
				//		if w.maxRetries == 1 {
				//			return
				//		}
				//
				//		time.Sleep(sleepDur)
				//		sleepDur *= 3
				//	}
				//}
				//

			case <-ctx.Done():
				return
			}
		}
	}(ctx, doneCh)
}

func (w *Worker[T]) Log(msg string, a ...any) {
	log.Println("[WORKER_POOL] " + fmt.Sprintf(msg, a...))
}
