package rlm

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type RateLimiter struct {
	rate  rate.Limit
	mutex *sync.RWMutex

	ipList     map[string]*Client
	bucketSize int

	params Params
}

type Params struct {
	MaxLimitRate  int
	BDIncrement   float64
	BlockDuration time.Duration
}
