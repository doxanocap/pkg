package rlm

import (
	"golang.org/x/time/rate"
	"time"
)

type Client struct {
	ip      string
	limiter *rate.Limiter

	blockDuration *time.Duration
	lastSeen      *time.Time
}

func newClient(limiter *rate.Limiter, ip string) *Client {
	return &Client{
		ip:            ip,
		limiter:       limiter,
		lastSeen:      ctxholder.GetPtr(time.Now()),
		blockDuration: ctxholder.GetPtr(time.Duration(0)),
	}
}

func (c *Client) BlockDuration() *time.Duration {
	diff := *c.blockDuration - time.Since(*c.lastSeen)
	if diff.Milliseconds() > 0 {
		return ctxholder.GetPtr(diff)
	}
	return nil
}

func (c *Client) HasLimit() bool {
	return c.blockDuration.Milliseconds() == 0
}

func (c *Client) Allow() bool {
	return c.limiter.Allow()
}

func (c *Client) AddRequest() {
	if !c.limiter.Allow() {

	}
}
