package rlm

import (
	"golang.org/x/time/rate"
	"time"
)

func (r *RateLimiter) GetByIP(ip string) *Client {
	client := r.getByIP(ip)
	if client == nil {
		return r.addLimiter(ip)
	}

	if !client.HasLimit() {
		blockDur := client.BlockDuration()
		if blockDur != nil {
			return client
		}
		client.lastSeen = nil
	}

	return client
}

func (r *RateLimiter) BlockByIP(ip string) {
	client := r.getByIP(ip)
	if client == nil {
		return
	}

	r.incrementBlock()
	client.lastSeen = ctxholder.GetPtr(time.Now())
	client.blockDuration = ctxholder.GetPtr(r.params.BlockDuration)
}

func (r *RateLimiter) addLimiter(ip string) *Client {
	client := newClient(rate.NewLimiter(r.rate, r.bucketSize), ip)
	r.set(ip, client)
	return client
}

func (r *RateLimiter) incrementBlock() {
	r.params.BlockDuration = time.Duration(r.params.BDIncrement * r.params.BlockDuration.Seconds())
}

func (r *RateLimiter) set(ip string, client *Client) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.ipList[ip] = client
}

func (r *RateLimiter) getByIP(ip string) *Client {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	v, ok := r.ipList[ip]
	if !ok {
		return nil
	}
	return v
}
