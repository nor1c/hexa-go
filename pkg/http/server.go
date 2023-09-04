package http

import (
	"context"
	"fmt"
	"gc-hexa-go/pkg/utils/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
)

type ServerOptions struct {
	Host string
	Port int
}

type HTTPServer struct {
	ServerOptions
}

func defaultOptions() ServerOptions {
	return ServerOptions{
		Host: "localhost",
		Port: 9000,
	}
}

func NewHTTPServer() *HTTPServer {
	opts := defaultOptions()

	return &HTTPServer{
		ServerOptions: opts,
	}
}

func gracefulShutdown(server *http.Server, ctx context.Context) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		logger.Logger.Warn().Msg("Shutting down gracefully..")

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Error during graceful shutdown: %s", err)
		}
	}()
}

func Serve() {
	r := chi.NewRouter()

	InitMiddlewares(r)
	InitRoutes(r)

	server := NewHTTPServer()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.Host, server.Port),
		Handler: r,
	}

	srvCtx, serverStopCtx := context.WithCancel(context.Background())
	gracefulShutdown(srv, srvCtx)

	logger.Logger.Info().Msg("Server started")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP ListenAndServe err: %v", err)
	}

	serverStopCtx()
}
