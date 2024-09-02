package workerpool

import (
	"context"
	"fmt"
	"github.com/doxanocap/pkg/workerpool/test_service"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	start := time.Now()
	tasks := test_service.GetTasks()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	wp := NewWorkerPool[test_service.Task, test_service.Result](
		10,
		test_service.HttpRequest)

	wp.Start(ctx)
	wp.AppendTasks(tasks)

	go func() {
		<-exit
		wp.Stop()
		cancel()
		fmt.Println("Gracefully shutdown WORKER_POOL test")
		os.Exit(0)
	}()

	defer func() {
		log.Println("number of goroutines before stop:", runtime.NumGoroutine())
		wp.Stop()
		cancel()
		log.Println("number of goroutines after stop:", runtime.NumGoroutine())
		log.Println("latency: ", time.Since(start))
	}()

	i := 0
	for r := range wp.GetResults() {
		if i == len(tasks)-1 {
			return
		}
		log.Printf("Data: %s | Err: %v \n", string(r.Data()), r.Err())
		i++
	}
}
