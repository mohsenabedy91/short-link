//go:build !test

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"short-link/internal/adapter/http/handler"
	"short-link/internal/adapter/http/routes"
	"short-link/internal/adapter/storage/postgres"
	"short-link/internal/adapter/storage/postgres/shortlinkrepository"
	"short-link/internal/core/config"
	"short-link/internal/core/service"
	"short-link/pkg/translation"
	"syscall"
	"time"

	_ "net/http/pprof"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
// @description "Bearer <your-jwt-token>"
func main() {
	conf := config.GetConfig()

	ctx := context.Background()
	defer func() {
		if err := postgres.Close(); err != nil {
			fmt.Printf("Close error: %v\n", err)
		}
	}()

	if err := postgres.InitClient(ctx, conf); err != nil {
		fmt.Printf("Initialize Database error: %v\n", err)
		return
	}

	postgresDB := postgres.Get()

	trans := translation.NewTranslation(conf.App)
	trans.GetLocalizer(conf.App.Locale)

	shortLinkRepo := shortlinkrepository.NewShortLinkRepository(postgresDB)

	shortLinkService := service.NewShortLinkService(conf.ShortLink, shortLinkRepo)
	shortLinkHandler := handler.NewShortLinkHandler(trans, shortLinkService)

	// Init router
	router, err := routes.NewRouter(conf, trans)
	if err != nil {
		fmt.Printf("New Router error: %v\n", err)
		return
	}

	router = router.NewShortLinkRouter(*shortLinkHandler)

	listenAddr := fmt.Sprintf("%s:%s", conf.App.URL, conf.App.Port)
	server := &http.Server{
		Addr:    listenAddr,
		Handler: router.Engine.Handler(),
	}

	// Start server
	router.Serve(server)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	timeout := conf.App.GracefullyShutdown * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		fmt.Printf("shutdown error: %v\n", err)
	}

	<-ctx.Done()
}
