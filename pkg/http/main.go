package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yashptel/go-api-template/pkg/controllers"
	"github.com/yashptel/go-api-template/pkg/middleware"
	"go.uber.org/zap"
)

func RunHttpServer() {
	router := controllers.NewRouter()

	router.Use(middleware.Logger)

	zap.L().Info("Starting http server")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Any("err", err))
		}
	}()
	zap.L().Sugar().Infof("Server started on port %s", srv.Addr)

	<-done
	zap.L().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: %s\n", zap.Any("err", err))
	}
	zap.L().Info("Server shutdown gracefully")
}
