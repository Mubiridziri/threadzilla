package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"threadzilla/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		fmt.Printf("Received signal: %s, shutting down...\n", sig)
		cancel()
	}()

	a := app.Application{}

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
