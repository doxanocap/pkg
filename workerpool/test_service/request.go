package test_service

import (
	"context"
	"fmt"
	"time"
)

func HttpRequest(ctx context.Context, task Task) Result {
	time.Sleep(100 * time.Millisecond)
	return Result{
		data: []byte(fmt.Sprintf("done %s", task.ID)),
		err:  nil,
	}
}
