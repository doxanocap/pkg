package internal

import (
	"golang.org/x/time/rate"
	"sync"
)

type Error struct {
	Status  int
	Message string
}

type Visitor struct {
	Limiter   *rate.Limiter
	LimitTime int64
	LastSeen  int64
}

type RateLimiter struct {
	IPList     map[string]*Visitor
	Mutex      *sync.RWMutex
	Rate       rate.Limit
	BucketSize int
}

type Params struct {
	MaxRate     int
	BlockTime   int64
	BTIncrement int
}
