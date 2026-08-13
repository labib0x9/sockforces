package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/labib0x9/sockforces/config"
	rest "github.com/labib0x9/sockforces/internal/transport/http"
)

func main() {
	cnf := config.GetConfig()

	fmt.Println(cnf.Service)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := rest.NewServer()

	go func() {
		server.Start(cnf)
	}()

	<-ctx.Done()

	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(shutdown)
}
