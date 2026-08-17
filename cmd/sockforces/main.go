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
	"github.com/labib0x9/sockforces/internal/infra/rabbitmq"
	rest "github.com/labib0x9/sockforces/internal/transport/http"
	subhandler "github.com/labib0x9/sockforces/internal/transport/http/handlers/submissions"
	"github.com/labib0x9/sockforces/internal/transport/http/middlewares"
	"github.com/labib0x9/sockforces/internal/worker"
)

func main() {
	cnf := config.GetConfig(".env")

	validate := validator.New()

	middleware := middlewares.NewMiddlewares(cnf)

	rmqClient := rabbitmq.NewClient(cnf.RabbitMQ)
	defer rmqClient.Close()

	if err := rmqClient.Setup(); err != nil {
		panic(err)
	}

	worker := worker.NewWorker(rmqClient)

	gitClient := github.NewClient(cnf.Github)
	gitRepo := github.NewGithubRepo(gitClient, cnf.Github)
	subService := subservice.NewService(gitRepo, rmqClient)
	subHandler := subhandler.NewHandler(subService, validate, middleware)

	server := rest.NewServer(*subHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go worker.Run(ctx, "test-container", 10)

	go func() {
		server.Start(cnf)
	}()

	<-ctx.Done()

	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(shutdown)
}
