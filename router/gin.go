package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kan-too/pkg/config"
	"net/http"
)

func InitGinRouter(env string) *gin.Engine {
	if env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.RedirectTrailingSlash = true
	corsConfig := cors.Config{
		//AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           604800,
	}
	router.Use(cors.New(corsConfig))
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	if env == config.EnvProduction {
		router.Use(gin.Logger())
	}
	return router
}
