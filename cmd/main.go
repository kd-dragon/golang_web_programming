package main

import (
	"GolangLivePT01/golang_web_programming/app/config"
	"GolangLivePT01/golang_web_programming/app/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title 	Membership API
// @version 1.0
// @host 	localhost:8080
func main() {

	fmt.Println(os.Getpid())

	cfg := config.GetInstance()
	routes.InitializeRoutes(cfg.GetGroup())

	go func() {
		if err := cfg.GetEcho().Start(":" + cfg.GetServicePort()); err != nil && err != http.ErrServerClosed {
			cfg.GetEcho().Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := cfg.GetEcho().Shutdown(ctx); err != nil {
		cfg.GetEcho().Logger.Fatal(err)
	}
}
