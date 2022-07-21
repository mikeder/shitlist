package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mikeder/shitlist/internal/config"
	"github.com/mikeder/shitlist/internal/server"
)

func main() {
	cfg := new(config.Specification)
	if err := cfg.LoadFromEnvironment(); err != nil {
		_ = cfg.Usage()
		log.Fatal(err)
	}

	srv := server.Setup(cfg)

	go func() {
		if err := server.Start(srv); err != nil {
			log.Fatal(err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	// NOTE: syscall.SIGKILL, syscall.SIGSTOP, and os.Kill cannot be trapped.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	// Waiting for signal to shutdown
	sig := <-stop
	log.Println("Shutdown due to " + sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
