package workerpool

import (
	"auto-close-contracts/internal/config"
	"auto-close-contracts/internal/logger"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	start := time.Now()
	cfg := config.InitConfig()

	log := logger.InitSlogLogger(cfg)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	tasks := NewTasks()
	wp := NewWorkerPool(10, func(ctx context.Context, t Task) error {
		_, err := HttpRequest(ctx, t)
		return err
	})
	wp.Start(ctx, tasks)

	go func() {
		<-exit
		cancel()
		fmt.Println("Gracefully shutdown WORKER_POOL test")
		os.Exit(0)
	}()

	defer func() {
		fmt.Printf("dur: %s\n", time.Since(start).String())
	}()

	countSuccessful := 0
	results := wp.GetResultsCh()
	for e := range results {
		if e == nil {
			countSuccessful++
		} else {
			log.Error(e.Error())
		}
	}

	fmt.Println(countSuccessful)
}

type Task struct {
	id  int
	val string
}

func (t Task) ID() uint64 {
	return uint64(t.id)
}

func HttpRequest(ctx context.Context, task Task) (int, error) {
	time.Sleep(1000 * time.Millisecond)
	n := rand.Intn(100)
	if n > 50 {
		return 0, fmt.Errorf("too many requests")
	}
	return n, nil
}

func NewTasks() []Task {
	tasks := make([]Task, 10)
	for i := range tasks {
		tasks[i].id = rand.Intn(10)
	}
	return tasks
}
