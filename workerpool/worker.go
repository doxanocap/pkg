package workerpool

import (
	"context"
	"sync"
)

type Worker[T ITask, R IResult] struct {
	tasks   <-chan T
	results chan<- R
	job     func(context.Context, T) R
	wg      *sync.WaitGroup
}

func SpanWorker[T ITask, R IResult](
	tasks <-chan T,
	results chan<- R,
	job func(ctx context.Context, task T) R,
	wg *sync.WaitGroup,
) *Worker[T, R] {
	return &Worker[T, R]{
		job:     job,
		tasks:   tasks,
		results: results,
		wg:      wg,
	}
}

func (w *Worker[T, R]) Start(ctx context.Context, doneCh chan<- struct{}) {
	go func(ctx context.Context, doneCh chan<- struct{}) {
		defer func() {
			doneCh <- struct{}{}
			w.wg.Done()
		}()

		w.wg.Add(1)
		for {
			select {
			case task, ok := <-w.tasks:
				if !ok {
					return
				}
				w.results <- w.job(ctx, task)
			case <-ctx.Done():
				return
			}
		}
	}(ctx, doneCh)
}
