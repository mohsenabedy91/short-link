package routes

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link/internal/adapter/http/middlewares"
	"short-link/internal/core/config"
	"short-link/pkg/translation"
)

// Router is a wrapper for HTTP router
type Router struct {
	Engine *gin.Engine
	conf   config.Config
	trans  translation.Translator
}

// NewRouter creates a new HTTP router
func NewRouter(conf config.Config, trans translation.Translator) (*Router, error) {

	// Disable debug mode in production
	if conf.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler(trans)))

	setSwaggerRoutes(router.Group(""), conf.Swagger)

	return &Router{
		Engine: router,
		conf:   conf,
		trans:  trans,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(server *http.Server) {
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error starting the HTTP server: %v", err)
		}
	}()
}
