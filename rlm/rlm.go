package rlm

import (
	"sync"
	"time"
)

func InitRLM(p Params) *RateLimiter {
	// default params
	if p.MaxLimitRate <= 0 {
		p.MaxLimitRate = 5
	}
	if p.BlockDuration.Milliseconds() <= 0 {
		p.BlockDuration = time.Minute * 2
	}
	if p.BDIncrement <= 1 {
		p.BDIncrement = 1.25
	}

	return &RateLimiter{
		mutex:      &sync.RWMutex{},
		ipList:     make(map[string]*Client),
		bucketSize: p.MaxLimitRate,
		rate:       1,
	}
}
