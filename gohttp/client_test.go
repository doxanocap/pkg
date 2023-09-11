package gohttp

import (
	"context"
	"fmt"
	"github.com/doxanocap/pkg/lg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(200, "pong")
	})
	return r
}

func Test(t *testing.T) {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lg.Fatalf("listen: %s\n", err)
		}
	}()

	res, err := NewRequest().
		SetURL("http://localhost:8080/ping").
		SetMethod(MethodGet).
		SetRequestFormat(FormatJSON).
		Execute(context.Background())
	if err != nil {
		lg.Fatalf("gohttp: %v", err)
	}

	fmt.Println(res.Status, res.Body)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	lg.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		lg.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		lg.Info("timeout of 5 seconds.")
	}
	lg.Info("Server exiting")
}
