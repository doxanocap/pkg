package workerpool

type Option[T ITask] func(wp *WorkerPool[T])

func NoRetries[T ITask]() Option[T] {
	return func(wp *WorkerPool[T]) {
		wp.maxRetries = 1
	}
}
