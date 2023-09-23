package rlm

import (
	"errors"
	"fmt"
	"github.com/doxanocap/pkg/rlm/internal"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRLM(MaxLimitRate, BlockTimeIncrement int, DefaultBlockTime int64) error {
	if MaxLimitRate <= 0 || DefaultBlockTime <= 0 || BlockTimeIncrement <= 0 {
		return errors.New("invalid params")
	}
	internal.DefaultParams = internal.Params{
		MaxRate:     MaxLimitRate,
		BlockTime:   DefaultBlockTime,
		BTIncrement: BlockTimeIncrement}
	return nil
}

// ToDo: complete package and write README.md
func RequestLimitMiddleware(ctx *gin.Context) {
	nilParams := internal.Params{}
	if internal.DefaultParams == nilParams {
		internal.DefaultParams = internal.Params{
			MaxRate:     5,
			BlockTime:   2,
			BTIncrement: 0,
		}
	}

	var mainLimiter = internal.InitRateLimiter(1, internal.DefaultParams.MaxRate)

	ipAddress := ctx.Request.RemoteAddr
	fwdAddress := ctx.ClientIP()
	// It parses IP from "X-Forwarded-For"
	if fwdAddress != "" {
		ipAddress = fwdAddress
	}

	l, dur := mainLimiter.GetLimiter(ipAddress)
	if dur != 0 {
		ctx.JSON(
			429,
			internal.Error{
				Status:  429,
				Message: fmt.Sprintf("Too many requests, you need to wait %s", time.Unix(dur-60*60*6, 0).Format("15:04:05")),
			})
		ctx.Abort()
		return
	}

	if !l.Allow() {
		mainLimiter.Lock(ipAddress, internal.DefaultParams.BlockTime)
		ctx.JSON(429, internal.Error{Status: 429, Message: "Too many requests"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
