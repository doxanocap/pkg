package gohttp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		value := c.GetHeader("TOKEN")
		if value == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		fmt.Println(value)
		c.String(http.StatusOK, value)
	})
	return r
}

func Test(t *testing.T) {
	r := setupRouter()
	ctx := context.Background()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	res, err := NewRequest().
		SetURL("http://localhost:8080/ping").
		SetMethod(http.MethodGet).
		SetHeader("TOKEN", "test123456").
		Execute(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("gohttp: %v", err))
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Server Shutdown: %s", err))
	}
}
