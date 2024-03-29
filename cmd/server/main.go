package main

import (
	"context"
	"flag"
	"fmt"
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		_ = cfg.Usage()
	}
	flag.Parse()

	if err := cfg.LoadFromEnvironment(); err != nil {
		_ = cfg.Usage()
		log.Fatal(err)
	}

	srv, err := server.Setup(cfg)
	if err != nil {
		log.Fatal(err)
	}

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
