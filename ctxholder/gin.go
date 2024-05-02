package ctxholder

import (
	"context"
	"sync"
)

const (
	ContextHolderKey = "context-holder-key"
	keyUserId        = "user_id"
	keyRefreshToken  = "refresh_token"
)

func SetKV(ctx context.Context, key string, value interface{}) {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		contextHolder.Store(key, value)
	}
}

func SetUserID(ctx context.Context, userID string) {
	SetKV(ctx, keyUserId, userID)
}

func GetUserID(ctx context.Context) string {
	value := getAttribute(ctx, keyUserId)
	if value != nil {
		if i, ok := value.(string); ok {
			return i
		}
	}
	return ""
}

func SetRefreshToken(ctx context.Context, token string) {
	SetKV(ctx, keyRefreshToken, token)
}

func GetRefreshToken(ctx context.Context) string {
	value := getAttribute(ctx, keyRefreshToken)
	if value != nil {
		if i, ok := value.(string); ok {
			return i
		}
	}
	return ""
}

func GetIntByKey(ctx context.Context, key string) int {
	value := getAttribute(ctx, key)
	if value != nil {
		if i, ok := value.(int); ok {
			return i
		}
	}
	return 0
}

func GetInt64ByKey(ctx context.Context, key string) int64 {
	value := getAttribute(ctx, key)
	if value != nil {
		if i, ok := value.(int64); ok {
			return i
		}
	}
	return 0
}

func GetStringByKey(ctx context.Context, key string) string {
	value := getAttribute(ctx, key)
	if value != nil {
		if i, ok := value.(string); ok {
			return i
		}
	}
	return ""
}

func getAttribute(ctx context.Context, key string) interface{} {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		value, ok := contextHolder.Load(key)
		if ok {
			return value
		}
	}
	return nil
}
