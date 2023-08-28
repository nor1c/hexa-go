package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

func gracefulShutdown(server *http.Server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown err: %v", err)
	}
}

func Serve() {
	InitRoutes()

	server := NewHTTPServer()

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", server.Host, server.Port),
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP ListenAndServe err: %v", err)
	}

	gracefulShutdown(srv)
}
