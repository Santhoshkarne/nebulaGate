package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SaisrikarVollala/nebulagate/internal/config"
	"github.com/SaisrikarVollala/nebulagate/internal/middleware"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NebulaGate - AI Powered Distributed Load Balancer"))
	})

	servers, err := config.LoadServers("config/servers.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range servers {
		fmt.Println(s.ID, s.URL)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Recovery(mux),
	}

	go func() {
		log.Println("Server Listening on :8080")
		if err := httpServer.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Listen for OS signals
	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	sig := <-sigChan

	log.Printf("received signal: %v", sig)
	log.Println("starting graceful shutdown...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Println("NebulaGate stopped gracefully")

}
