package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anshulbhargav1/student-api/Internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student-api"))
	})

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started ", slog.String("address", cfg.Addr))

	// Channel to listen for interrupt or termination signal (Ctrl+C)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	// Run server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to server")
		}
	}()

	// Block until we receive a signal (Ctrl+C)
	<-done
	slog.Info("Shutting down the server")

	// Create a deadline (e.g., 5 seconds) for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Forced to shutdown: %s", err)
	}

	slog.Info("Server shutdown succesfully..")

}
