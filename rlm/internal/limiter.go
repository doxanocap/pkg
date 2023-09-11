package internal

import (
	"golang.org/x/time/rate"
	"time"
)

func (r *RateLimiter) AddLimit(ip string) *rate.Limiter {
	r.Mutex.Lock()
	limiter := rate.NewLimiter(r.Rate, r.BucketSize)
	r.IPList[ip] = &Visitor{
		limiter,
		0,
		0}
	r.Mutex.Unlock()
	return limiter
}

func (r *RateLimiter) GetLimiter(ip string) (*rate.Limiter, int64) {
	r.Mutex.Lock()
	visitor, ok := r.IPList[ip]

	if !ok {
		r.Mutex.Unlock()
		return r.AddLimit(ip), 0
	}

	if visitor.LimitTime != 0 {
		diff := visitor.LimitTime - (time.Now().Unix() - visitor.LastSeen)
		if diff > 0 {
			r.Mutex.Unlock()
			return nil, diff
		} else {
			r.UnLock(ip)
		}
	}

	r.Mutex.Unlock()
	return visitor.Limiter, 0
}

func (r *RateLimiter) Lock(ip string, dur int64) {
	r.IPList[ip].LastSeen = time.Now().Unix()
	r.IPList[ip].LimitTime = dur
}

func (r *RateLimiter) UnLock(ip string) {
	r.IPList[ip].LastSeen = 0
}
