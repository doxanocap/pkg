package ctxholder

import (
	"github.com/gin-gonic/gin"
)

const (
	keyUserId       = "user_id"
	keyRefreshToken = "refresh_token"
	keyClientIP     = "client_ip"
)

func SetClientIP(c *gin.Context) {
	c.Set(keyClientIP, c.ClientIP())
}

func SetString(c *gin.Context, key, value string) {
	c.Set(key, value)
}

func SetUserID(c *gin.Context, userID string) {
	c.Set(keyUserId, userID)
}

func GetUserID(c *gin.Context) string {
	v, ok := c.Get(keyUserId)
	if !ok {
		return ""
	}
	return v.(string)
}

func SetRefreshToken(c *gin.Context, token string, age int) {
	c.SetCookie(keyRefreshToken,
		token,
		int(age),
		"/",
		"localhost",
		false,
		true)
	c.Set(keyRefreshToken, token)
}

func GetRefreshToken(c *gin.Context) string {
	v, ok := c.Get(keyRefreshToken)
	if !ok {
		return ""
	}
	return v.(string)
}

func GetClientIP(c *gin.Context) string {
	v, ok := c.Get(keyClientIP)
	if !ok {
		return ""
	}
	return v.(string)
}

func GetIntByKey(c *gin.Context, key string) int {
	v, ok := c.Get(key)
	if !ok {
		return 0
	}
	return v.(int)
}

func GetString(c *gin.Context, key string) string {
	v, ok := c.Get(key)
	if !ok {
		return ""
	}
	return v.(string)
}
