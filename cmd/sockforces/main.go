package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/sockforces/config"
	subservice "github.com/labib0x9/sockforces/internal/app/submissions"
	"github.com/labib0x9/sockforces/internal/infra/github"
	rest "github.com/labib0x9/sockforces/internal/transport/http"
	subhandler "github.com/labib0x9/sockforces/internal/transport/http/handlers/submissions"
	"github.com/labib0x9/sockforces/internal/transport/http/middlewares"
)

func main() {
	cnf := config.GetConfig(".env")

	validate := validator.New()

	middleware := middlewares.NewMiddlewares(cnf)

	gitClient := github.NewClient(cnf.Github)
	gitRepo := github.NewGithubRepo(gitClient, cnf.Github)
	subService := subservice.NewService(gitRepo)
	subHandler := subhandler.NewHandler(subService, validate, middleware)

	server := rest.NewServer(*subHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		server.Start(cnf)
	}()

	<-ctx.Done()

	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(shutdown)
}
