package internal

import (
	"golang.org/x/time/rate"
	"sync"
)

var DefaultParams Params

func InitRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		IPList:     make(map[string]*Visitor),
		Mutex:      &sync.RWMutex{},
		Rate:       r,
		BucketSize: b,
	}
}
