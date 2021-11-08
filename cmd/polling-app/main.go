package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/milessh/polling-app/internal/backend"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c

		log.Println("System call: ", oscall)

		cancel()
	}()

	start(ctx)
}

func start(ctx context.Context) {
	go backend.Start()

	log.Println("server started")

	// wait for interrupt channel end signal
	<-ctx.Done()

	backend.Terminate()

	log.Println("server stopped")
}
