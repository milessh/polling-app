package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/milessh/polling-app/pkg/router"
)

var s *http.Server

func Init() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/", router.Serve)

	s = &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
}

func Start() {
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start: ", err)
	}
}

func Terminate() (err error) {
	log.Println("Terminating server")

	ctxEnd, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = s.Shutdown(ctxEnd); err != nil {
		log.Fatal("Failed to terminate: ", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	log.Println("Server terminated")

	return
}
