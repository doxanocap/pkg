package rlm

import (
	"fmt"
	"github.com/doxanocap/pkg/sandbox/lg"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rlm := InitRLM(Params{
		MaxLimitRate:  10,
		BDIncrement:   1,
		BlockDuration: time.Minute * 2,
	})

	ip := "192.3.45.123"

	client := rlm.GetByIP(ip)

	if duration := client.BlockDuration(); duration != nil {
		lg.Error(fmt.Errorf("blocked for: %f minutes", duration.Minutes()))
		return
	}

	if !client.Allow() {
		rlm.BlockByIP(ip)
	}
}
